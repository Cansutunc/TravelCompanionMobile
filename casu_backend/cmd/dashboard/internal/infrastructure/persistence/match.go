package persistence

import (
	"context"
	"time"
)

// Match the client persistence model interface
type Match interface {
	GetMatchID() string
	GetCompetitionID() string
	GetPlayingTime() time.Time
	GetAwayTeamID() string
	GetHomeTeamID() string
	GetSubDescription() string
	GetYear() string
	GetAreaID() string
}

// MatchRepository allows to get/save current state of user to memory storage
type MatchRepository interface {
	Get(ctx context.Context, id string) (Match, error)
	Add(ctx context.Context, match Match) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, id string) (Match, error)
	GetByName(ctx context.Context, name string) (Match, error)
	FindAll(ctx context.Context, limit, offset int64) ([]Match, error)
	Update(ctx context.Context, id, name string, code string, matchId string) error
}
