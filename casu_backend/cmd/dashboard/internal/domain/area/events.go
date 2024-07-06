package area

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
func (e WasCreated) GetAreaID() string {
	return e.AreaID.String()
}

// GetName the name
func (e WasCreated) GetName() string {
	return e.Name
}

// GetCountryID the countryID
func (e WasCreated) GetCityID() string {
	return e.CityID.String()
}

// WasCreated event
type WasCreated struct {
	AreaID uuid.UUID `json:"area_id" bson:"area_id"`
	CityID uuid.UUID `json:"city_id" bson:"city_id"`
	Name   string    `json:"name" bson:"name"`
}

// WasRemoved event
type WasRemoved struct {
	AreaID uuid.UUID `json:"area_id" bson:"area_id"`
}

// WasUpdated event
type WasUpdated struct {
	AreaID uuid.UUID `json:"area_id" bson:"area_id"`
	Name   string    `json:"name,omitempty" bson:"name,omitempty"`
	CityID uuid.UUID `json:"city_id,omitempty" bson:"city_id,omitempty"`
}
