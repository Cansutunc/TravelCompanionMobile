package persistence

import "context"

// Country the client persistence model interface
type Area interface {
	GetAreaID() string
	GetCityID() string
	GetName() string
}

// CountryRepository allows to get/save current state of user to memory storage
type AreaRepository interface {
	Get(ctx context.Context, id string) (Area, error)
	Add(ctx context.Context, client Area) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, id string) (Area, error)
	GetByName(ctx context.Context, name string) (Area, error)
	FindAll(ctx context.Context, limit, offset int64) ([]Area, error)
	Update(ctx context.Context, id, name string, cityID string) error
}
