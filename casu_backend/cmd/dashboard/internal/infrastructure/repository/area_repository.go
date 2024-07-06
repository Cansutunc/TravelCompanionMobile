/*
Package repository holds event sourced repositories
*/
package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/area"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/eventbus"
	"github.com/hsynrtn/dashboard-management/pkg/eventstore"
)

type areaRepository struct {
	eventStore eventstore.EventStore
	eventBus   eventbus.EventBus
}

// Save current area changes to event store and publish each event with an event bus
func (r *areaRepository) Save(ctx context.Context, u area.Area) error {
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

// Get area with current state applied
func (r *areaRepository) Get(ctx context.Context, id uuid.UUID) (area.Area, error) {
	events, err := r.eventStore.GetStream(ctx, id, area.StreamName)
	if err != nil {
		return area.Area{}, apperrors.Wrap(err)
	}

	if len(events) == 0 {
		return area.Area{}, apperrors.ErrNotFound
	}

	return area.FromHistory(ctx, events)
}

// NewTokenRepository creates new area event sourced repository
func NewAreaRepository(store eventstore.EventStore, bus eventbus.EventBus) area.Repository {
	return &areaRepository{store, bus}
}
