package country

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

func (e WasCreated) GetCountryID() string {
	return e.CountryID.String()
}
func (e WasCreated) GetProvince() string {
	return e.Province
}
func (e WasCreated) GetCountry() string {
	return e.Country
}
func (e WasCreated) GetUserId() string {
	return e.UserId
}
func (e WasCreated) GetUserName() string {
	return e.UserName
}

// WasRemoved event
type WasRemoved struct {
	CountryID uuid.UUID `json:"country_id" bson:"country_id"`
}

type WasCreated struct {
	CountryID uuid.UUID `json:"country_id" bson:"country_id"`
	UserId    string    `json:"user_id" bson:"user_id"`
	Province  string    `json:"province" bson:"province"`
	Country   string    `json:"country" bson:"country"`
	UserName  string    `json:"user_name" bson:"user_name"`
}

type WasUpdated struct {
	CountryID uuid.UUID `json:"country_id" bson:"country_id"`
	Province  string    `json:"province,omitempty" bson:"province,omitempty"`
	Country   string    `json:"country,omitempty" bson:"country,omitempty"`
	UserId    string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	UserName  string    `json:"user_name,omitempty" bson:"user_name,omitempty"`
}
