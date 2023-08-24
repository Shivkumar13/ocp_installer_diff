package models
import (
    "errors"
)
// Provides operations to manage the collection of agreement entities.
type InstallState int

const (
    // Not Applicable.
    NOTAPPLICABLE_INSTALLSTATE InstallState = iota
    // Installed.
    INSTALLED_INSTALLSTATE
    // Failed.
    FAILED_INSTALLSTATE
    // Not Installed.
    NOTINSTALLED_INSTALLSTATE
    // Uninstall Failed.
    UNINSTALLFAILED_INSTALLSTATE
    // Unknown.
    UNKNOWN_INSTALLSTATE
)

func (i InstallState) String() string {
    return []string{"notApplicable", "installed", "failed", "notInstalled", "uninstallFailed", "unknown"}[i]
}
func ParseInstallState(v string) (interface{}, error) {
    result := NOTAPPLICABLE_INSTALLSTATE
    switch v {
        case "notApplicable":
            result = NOTAPPLICABLE_INSTALLSTATE
        case "installed":
            result = INSTALLED_INSTALLSTATE
        case "failed":
            result = FAILED_INSTALLSTATE
        case "notInstalled":
            result = NOTINSTALLED_INSTALLSTATE
        case "uninstallFailed":
            result = UNINSTALLFAILED_INSTALLSTATE
        case "unknown":
            result = UNKNOWN_INSTALLSTATE
        default:
            return 0, errors.New("Unknown InstallState value: " + v)
    }
    return &result, nil
}
func SerializeInstallState(values []InstallState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
