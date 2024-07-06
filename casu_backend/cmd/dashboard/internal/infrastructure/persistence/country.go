package persistence

import "context"

// Country the client persistence model interface
type Country interface {
	GetCountryID() string
	GetProvince() string
	GetCountry() string
	GetUserId() string
	GetUserName() string
}

// CountryRepository allows to get/save current state of user to memory storage
type CountryRepository interface {
	Get(ctx context.Context, id string) (Country, error)
	Add(ctx context.Context, client Country) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, id string) (Country, error)
	GetByName(ctx context.Context, name string) (Country, error)
	GetByUserIdAndProvince(ctx context.Context, userId string, province string) (Country, error)
	FindAll(ctx context.Context, limit, offset int64) ([]Country, error)
	Update(ctx context.Context, id, country string, province string, userId string) error
}
