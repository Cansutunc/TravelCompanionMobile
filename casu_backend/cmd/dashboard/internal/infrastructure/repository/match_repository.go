/*
Package repository holds event sourced repositories
*/
package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/match"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/eventbus"
	"github.com/hsynrtn/dashboard-management/pkg/eventstore"
)

type matchRepository struct {
	eventStore eventstore.EventStore
	eventBus   eventbus.EventBus
}

// Save current match changes to event store and publish each event with an event bus
func (r *matchRepository) Save(ctx context.Context, u match.Match) error {
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

// Get match with current state applied
func (r *matchRepository) Get(ctx context.Context, id uuid.UUID) (match.Match, error) {
	events, err := r.eventStore.GetStream(ctx, id, match.StreamName)
	if err != nil {
		return match.Match{}, apperrors.Wrap(err)
	}

	if len(events) == 0 {
		return match.Match{}, apperrors.ErrNotFound
	}

	return match.FromHistory(ctx, events)
}

// NewTokenRepository creates new match event sourced repository
func NewMatchRepository(store eventstore.EventStore, bus eventbus.EventBus) match.Repository {
	return &matchRepository{store, bus}
}
