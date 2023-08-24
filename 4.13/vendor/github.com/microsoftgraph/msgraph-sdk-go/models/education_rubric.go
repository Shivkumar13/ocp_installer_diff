package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationRubric 
type EducationRubric struct {
    Entity
    // The user who created this resource.
    createdBy IdentitySetable
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description of this rubric.
    description EducationItemBodyable
    // The name of this rubric.
    displayName *string
    // The grading type of this rubric -- null for a no-points rubric, or educationAssignmentPointsGradeType for a points rubric.
    grading EducationAssignmentGradeTypeable
    // The last user to modify the resource.
    lastModifiedBy IdentitySetable
    // Moment in time when the resource was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The collection of levels making up this rubric.
    levels []RubricLevelable
    // The collection of qualities making up this rubric.
    qualities []RubricQualityable
}
// NewEducationRubric instantiates a new educationRubric and sets the default values.
func NewEducationRubric()(*EducationRubric) {
    m := &EducationRubric{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationRubricFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationRubricFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationRubric(), nil
}
// GetCreatedBy gets the createdBy property value. The user who created this resource.
func (m *EducationRubric) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationRubric) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description of this rubric.
func (m *EducationRubric) GetDescription()(EducationItemBodyable) {
    return m.description
}
// GetDisplayName gets the displayName property value. The name of this rubric.
func (m *EducationRubric) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationRubric) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentitySetFromDiscriminatorValue , m.SetCreatedBy)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEducationItemBodyFromDiscriminatorValue , m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["grading"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEducationAssignmentGradeTypeFromDiscriminatorValue , m.SetGrading)
    res["lastModifiedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentitySetFromDiscriminatorValue , m.SetLastModifiedBy)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["levels"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRubricLevelFromDiscriminatorValue , m.SetLevels)
    res["qualities"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRubricQualityFromDiscriminatorValue , m.SetQualities)
    return res
}
// GetGrading gets the grading property value. The grading type of this rubric -- null for a no-points rubric, or educationAssignmentPointsGradeType for a points rubric.
func (m *EducationRubric) GetGrading()(EducationAssignmentGradeTypeable) {
    return m.grading
}
// GetLastModifiedBy gets the lastModifiedBy property value. The last user to modify the resource.
func (m *EducationRubric) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Moment in time when the resource was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationRubric) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLevels gets the levels property value. The collection of levels making up this rubric.
func (m *EducationRubric) GetLevels()([]RubricLevelable) {
    return m.levels
}
// GetQualities gets the qualities property value. The collection of qualities making up this rubric.
func (m *EducationRubric) GetQualities()([]RubricQualityable) {
    return m.qualities
}
// Serialize serializes information the current object
func (m *EducationRubric) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("grading", m.GetGrading())
        if err != nil {
            return err
        }
    }
    if m.GetLevels() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLevels())
        err = writer.WriteCollectionOfObjectValues("levels", cast)
        if err != nil {
            return err
        }
    }
    if m.GetQualities() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetQualities())
        err = writer.WriteCollectionOfObjectValues("qualities", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedBy sets the createdBy property value. The user who created this resource.
func (m *EducationRubric) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationRubric) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description of this rubric.
func (m *EducationRubric) SetDescription(value EducationItemBodyable)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name of this rubric.
func (m *EducationRubric) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGrading sets the grading property value. The grading type of this rubric -- null for a no-points rubric, or educationAssignmentPointsGradeType for a points rubric.
func (m *EducationRubric) SetGrading(value EducationAssignmentGradeTypeable)() {
    m.grading = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The last user to modify the resource.
func (m *EducationRubric) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Moment in time when the resource was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationRubric) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLevels sets the levels property value. The collection of levels making up this rubric.
func (m *EducationRubric) SetLevels(value []RubricLevelable)() {
    m.levels = value
}
// SetQualities sets the qualities property value. The collection of qualities making up this rubric.
func (m *EducationRubric) SetQualities(value []RubricQualityable)() {
    m.qualities = value
}
