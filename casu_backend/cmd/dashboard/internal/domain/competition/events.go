package competition

import (
	"fmt"
	"github.com/google/uuid"
)

var (
	WasCreatedType = (WasCreated{}).GetType()
	WasRemovedType = (WasRemoved{}).GetType()
	WasUpdatedType = (WasUpdated{}).GetType()
)

// GetType returns event type
func (e WasCreated) GetType() string {
	return fmt.Sprintf("%T", e)
}
func (e WasRemoved) GetType() string {
	return fmt.Sprintf("%T", e)
}
func (e WasUpdated) GetType() string {
	return fmt.Sprintf("%T", e)
}

// GetID the id
func (e WasCreated) GetCompetitionID() string {
	return e.CompetitionID.String()
}

// GetName the name
func (e WasCreated) GetName() string {
	return e.Name
}

// GetCode the code
func (e WasCreated) GetCompetitionType() string {
	return e.CompetitionType
}

// GetCountryID the countryID
func (e WasCreated) GetRegion() string {
	return e.Region
}

// GetCountryID the countryID
func (e WasCreated) GetSporType() string {
	return e.SporType
}

// WasCreated event
type WasCreated struct {
	CompetitionID   uuid.UUID `json:"competition_id" bson:"competition_id"`
	Name            string    `json:"name" bson:"name"`
	SporType        string    `json:"spor_type" bson:"spor_type"`
	Region          string    `json:"region" bson:"region"`
	CompetitionType string    `json:"competition_type" bson:"competition_type"`
}

// WasRemoved event
type WasRemoved struct {
	CompetitionID uuid.UUID `json:"competition_id" bson:"competition_id"`
}

// WasUpdated event
type WasUpdated struct {
	CompetitionID uuid.UUID `json:"competition_id" bson:"competition_id"`
	Name          string    `json:"name,omitempty" bson:"name,omitempty"`
	Code          string    `json:"code,omitempty" bson:"code,omitempty"`
	CountryID     uuid.UUID `json:"country_id,omitempty" bson:"country_id,omitempty"`
}
