/*
Package repository holds event sourced repositories
*/
package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/competition"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/eventbus"
	"github.com/hsynrtn/dashboard-management/pkg/eventstore"
)

type competitionRepository struct {
	eventStore eventstore.EventStore
	eventBus   eventbus.EventBus
}

// Save current competition changes to event store and publish each event with an event bus
func (r *competitionRepository) Save(ctx context.Context, u competition.Competition) error {
	if err := r.eventStore.Store(ctx, u.Changes()); err != nil {
		return apperrors.Wrap(err)
	}

	for _, event := range u.Changes() {
		if err := r.eventBus.Publish(ctx, event); err != nil {
			return apperrors.Wrap(err)
		}
	}

	return nil
}

// Get competition with current state applied
func (r *competitionRepository) Get(ctx context.Context, id uuid.UUID) (competition.Competition, error) {
	events, err := r.eventStore.GetStream(ctx, id, competition.StreamName)
	if err != nil {
		return competition.Competition{}, apperrors.Wrap(err)
	}

	if len(events) == 0 {
		return competition.Competition{}, apperrors.ErrNotFound
	}

	return competition.FromHistory(ctx, events)
}

// NewTokenRepository creates new competition event sourced repository
func NewCompetitionRepository(store eventstore.EventStore, bus eventbus.EventBus) competition.Repository {
	return &competitionRepository{store, bus}
}
