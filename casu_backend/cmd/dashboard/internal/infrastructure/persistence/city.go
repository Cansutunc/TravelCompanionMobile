package persistence

import "context"

// City the client persistence model interface
type City interface {
	GetCityID() string
	GetTitle() string
	GetDescription() string
	GetMaxPerson() int
	GetCreatedUserId() string
	GetCreatedUsername() string
	GetAllPerson() []string
}

// CityRepository allows to get/save current state of user to memory storage
type CityRepository interface {
	Get(ctx context.Context, id string) (City, error)
	Add(ctx context.Context, city City) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, id string) (City, error)
	GetByName(ctx context.Context, name string) (City, error)
	FindAll(ctx context.Context, limit, offset int64) ([]City, error)
	FindAllAccepted(ctx context.Context, username string) ([]City, error)
	Update(ctx context.Context, id, name string, code string, cityId string) error
	Accept(ctx context.Context, id, name string) error
	Remove(ctx context.Context, id, name string) error
}
