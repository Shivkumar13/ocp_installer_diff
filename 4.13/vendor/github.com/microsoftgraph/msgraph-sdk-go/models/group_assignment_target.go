package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupAssignmentTarget 
type GroupAssignmentTarget struct {
    DeviceAndAppManagementAssignmentTarget
    // The group Id that is the target of the assignment.
    groupId *string
}
// NewGroupAssignmentTarget instantiates a new GroupAssignmentTarget and sets the default values.
func NewGroupAssignmentTarget()(*GroupAssignmentTarget) {
    m := &GroupAssignmentTarget{
        DeviceAndAppManagementAssignmentTarget: *NewDeviceAndAppManagementAssignmentTarget(),
    }
    odataTypeValue := "#microsoft.graph.groupAssignmentTarget";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupAssignmentTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupAssignmentTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.exclusionGroupAssignmentTarget":
                        return NewExclusionGroupAssignmentTarget(), nil
                }
            }
        }
    }
    return NewGroupAssignmentTarget(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupAssignmentTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceAndAppManagementAssignmentTarget.GetFieldDeserializers()
    res["groupId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetGroupId)
    return res
}
// GetGroupId gets the groupId property value. The group Id that is the target of the assignment.
func (m *GroupAssignmentTarget) GetGroupId()(*string) {
    return m.groupId
}
// Serialize serializes information the current object
func (m *GroupAssignmentTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceAndAppManagementAssignmentTarget.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupId sets the groupId property value. The group Id that is the target of the assignment.
func (m *GroupAssignmentTarget) SetGroupId(value *string)() {
    m.groupId = value
}
