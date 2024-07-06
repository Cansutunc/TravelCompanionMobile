package persistence

import "context"

// Competition the client persistence model interface
type Competition interface {
	GetCompetitionID() string
	GetName() string
	GetRegion() string
	GetSporType() string
	GetCompetitionType() string
}

// CompetitionRepository allows to get/save current state of user to memory storage
type CompetitionRepository interface {
	Get(ctx context.Context, id string) (Competition, error)
	Add(ctx context.Context, competition Competition) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, id string) (Competition, error)
	GetByName(ctx context.Context, name string) (Competition, error)
	FindAll(ctx context.Context, limit, offset int64) ([]Competition, error)
	Update(ctx context.Context, id, name string, code string, competitionId string) error
}
