package match

import (
	"context"
	"github.com/google/uuid"
)

// Repository allows to get/save events from/to event store
type Repository interface {
	Save(ctx context.Context, c Match) error
	Get(ctx context.Context, id uuid.UUID) (Match, error)
}
