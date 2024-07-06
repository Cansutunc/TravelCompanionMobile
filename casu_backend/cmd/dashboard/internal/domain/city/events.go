package city

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
func (e WasCreated) GetCityID() string {
	return e.CityID.String()
}

// GetName the name
func (e WasCreated) GetTitle() string {
	return e.Title
}

// GetCode the code
func (e WasCreated) GetDescription() string {
	return e.Description
}

// GetCountryID the countryID
func (e WasCreated) GetMaxPerson() int {
	return e.MaxPerson
}

func (e WasCreated) GetCreatedUserId() string {
	return e.CreatedUserId
}

func (e WasCreated) GetCreatedUsername() string {
	return e.CreatedUsername
}

func (e WasCreated) GetAllPerson() []string {
	return e.AllPerson
}

// WasCreated event
type WasCreated struct {
	CityID          uuid.UUID `json:"city_id" bson:"city_id"`
	Title           string    `json:"title" bson:"title"`
	Description     string    `json:"description" bson:"description"`
	MaxPerson       int       `json:"max_person" bson:"max_person"`
	CreatedUserId   string    `json:"created_user_id" bson:"created_user_id"`
	CreatedUsername string    `json:"created_username" bson:"created_username"`
	AllPerson       []string  `json:"all_person" bson:"all_person"`
}

// WasRemoved event
type WasRemoved struct {
	CityID uuid.UUID `json:"city_id" bson:"city_id"`
}

// WasUpdated event
type WasUpdated struct {
	CityID    uuid.UUID `json:"city_id" bson:"city_id"`
	Name      string    `json:"name,omitempty" bson:"name,omitempty"`
	Code      string    `json:"code,omitempty" bson:"code,omitempty"`
	CountryID uuid.UUID `json:"country_id,omitempty" bson:"country_id,omitempty"`
}
