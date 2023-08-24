package cluster

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/cluster/aws"
	"github.com/openshift/installer/pkg/asset/cluster/azure"
	"github.com/openshift/installer/pkg/asset/cluster/openstack"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/password"
	"github.com/openshift/installer/pkg/asset/quota"
	"github.com/openshift/installer/pkg/metrics/timer"
	"github.com/openshift/installer/pkg/terraform"
	platformstages "github.com/openshift/installer/pkg/terraform/stages/platform"
	typesaws "github.com/openshift/installer/pkg/types/aws"
	typesazure "github.com/openshift/installer/pkg/types/azure"
	typesopenstack "github.com/openshift/installer/pkg/types/openstack"
	typesvsphere "github.com/openshift/installer/pkg/types/vsphere"
)

var (
	// InstallDir is the directory containing install assets.
	InstallDir string
)

// Cluster uses the terraform executable to launch a cluster
// with the given terraform tfvar and generated templates.
type Cluster struct {
	FileList []*asset.File
}

var _ asset.WritableAsset = (*Cluster)(nil)

// Name returns the human-friendly name of the asset.
func (c *Cluster) Name() string {
	return "Cluster"
}

// Dependencies returns the direct dependency for launching
// the cluster.
func (c *Cluster) Dependencies() []asset.Asset {
	return []asset.Asset{
		&installconfig.ClusterID{},
		&installconfig.InstallConfig{},
		// PlatformCredsCheck, PlatformPermsCheck and PlatformProvisionCheck
		// perform validations & check perms required to provision infrastructure.
		// We do not actually use them in this asset directly, hence
		// they are put in the dependencies but not fetched in Generate.
		&installconfig.PlatformCredsCheck{},
		&installconfig.PlatformPermsCheck{},
		&installconfig.PlatformProvisionCheck{},
		&quota.PlatformQuotaCheck{},
		&TerraformVariables{},
		&password.KubeadminPassword{},
	}
}

// Generate launches the cluster and generates the terraform state file on disk.
func (c *Cluster) Generate(parents asset.Parents) (err error) {
	if InstallDir == "" {
		logrus.Fatalf("InstallDir has not been set for the %q asset", c.Name())
	}

	clusterID := &installconfig.ClusterID{}
	installConfig := &installconfig.InstallConfig{}
	terraformVariables := &TerraformVariables{}
	parents.Get(clusterID, installConfig, terraformVariables)

	if fs := installConfig.Config.FeatureSet; strings.HasSuffix(string(fs), "NoUpgrade") {
		logrus.Warnf("FeatureSet %q is enabled. This FeatureSet does not allow upgrades and may affect the supportability of the cluster.", fs)
	}

	if installConfig.Config.Platform.None != nil {
		return errors.New("cluster cannot be created with platform set to 'none'")
	}

	if installConfig.Config.BootstrapInPlace != nil {
		return errors.New("cluster cannot be created with bootstrapInPlace set")
	}

	platform := installConfig.Config.Platform.Name()

	if azure := installConfig.Config.Platform.Azure; azure != nil && azure.CloudName == typesazure.StackCloud {
		platform = typesazure.StackTerraformName
	}

	if vsphere := installConfig.Config.Platform.VSphere; vsphere != nil {
		if len(vsphere.FailureDomains) != 0 {
			platform = typesvsphere.ZoningTerraformName
		}
	}

	stages := platformstages.StagesForPlatform(platform)

	terraformDir := filepath.Join(InstallDir, "terraform")
	if err := os.Mkdir(terraformDir, 0777); err != nil {
		return errors.Wrap(err, "could not create the terraform directory")
	}

	terraformDirPath, err := filepath.Abs(terraformDir)
	if err != nil {
		return errors.Wrap(err, "cannot get absolute path of terraform directory")
	}

	defer os.RemoveAll(terraformDir)
	terraform.UnpackTerraform(terraformDirPath, stages)

	logrus.Infof("Creating infrastructure resources...")
	switch platform {
	case typesaws.Name:
		if err := aws.PreTerraform(context.TODO(), clusterID.InfraID, installConfig); err != nil {
			return err
		}
	case typesazure.Name, typesazure.StackTerraformName:
		if err := azure.PreTerraform(context.TODO(), clusterID.InfraID, installConfig); err != nil {
			return err
		}
	case typesopenstack.Name:
		if err := openstack.PreTerraform(context.TODO(), clusterID.InfraID, installConfig); err != nil {
			return err
		}
	}

	tfvarsFiles := make([]*asset.File, 0, len(terraformVariables.Files())+len(stages))
	for _, file := range terraformVariables.Files() {
		tfvarsFiles = append(tfvarsFiles, file)
	}

	for _, stage := range stages {
		outputs, err := c.applyStage(platform, stage, terraformDirPath, tfvarsFiles)
		if err != nil {
			return errors.Wrapf(err, "failure applying terraform for %q stage", stage.Name())
		}
		tfvarsFiles = append(tfvarsFiles, outputs)
		c.FileList = append(c.FileList, outputs)
	}

	return nil
}

// Files returns the FileList generated by the asset.
func (c *Cluster) Files() []*asset.File {
	return c.FileList
}

// Load returns error if the tfstate file is already on-disk, because we want to
// prevent user from accidentally re-launching the cluster.
func (c *Cluster) Load(f asset.FileFetcher) (found bool, err error) {
	matches, err := filepath.Glob("terraform(.*)?.tfstate")
	if err != nil {
		return true, err
	}
	if len(matches) != 0 {
		return true, errors.Errorf("terraform state files alread exist.  There may already be a running cluster")
	}

	return false, nil
}

func (c *Cluster) applyStage(platform string, stage terraform.Stage, terraformDir string, tfvarsFiles []*asset.File) (*asset.File, error) {
	// Copy the terraform.tfvars to a temp directory which will contain the terraform plan.
	tmpDir, err := ioutil.TempDir("", fmt.Sprintf("openshift-install-%s-", stage.Name()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create temp dir for terraform execution")
	}
	defer os.RemoveAll(tmpDir)

	var extraOpts []tfexec.ApplyOption
	for _, file := range tfvarsFiles {
		if err := ioutil.WriteFile(filepath.Join(tmpDir, file.Filename), file.Data, 0600); err != nil {
			return nil, err
		}
		extraOpts = append(extraOpts, tfexec.VarFile(filepath.Join(tmpDir, file.Filename)))
	}

	return c.applyTerraform(tmpDir, platform, stage, terraformDir, extraOpts...)
}

func (c *Cluster) applyTerraform(tmpDir string, platform string, stage terraform.Stage, terraformDir string, opts ...tfexec.ApplyOption) (*asset.File, error) {
	timer.StartTimer(stage.Name())
	defer timer.StopTimer(stage.Name())

	applyErr := terraform.Apply(tmpDir, platform, stage, terraformDir, opts...)

	// Write the state file to the install directory even if the apply failed.
	if data, err := ioutil.ReadFile(filepath.Join(tmpDir, terraform.StateFilename)); err == nil {
		c.FileList = append(c.FileList, &asset.File{
			Filename: stage.StateFilename(),
			Data:     data,
		})
	} else if !os.IsNotExist(err) {
		logrus.Errorf("Failed to read tfstate: %v", err)
		return nil, errors.Wrap(err, "failed to read tfstate")
	}

	if applyErr != nil {
		return nil, errors.Wrap(applyErr, asset.ClusterCreationError)
	}

	outputs, err := terraform.Outputs(tmpDir, terraformDir)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get outputs from stage %q", stage.Name())
	}

	outputsFile := &asset.File{
		Filename: stage.OutputsFilename(),
		Data:     outputs,
	}
	return outputsFile, nil
}
