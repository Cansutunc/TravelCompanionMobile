package area

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hsynrtn/dashboard-management/pkg/domain"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/identity"
	"github.com/hsynrtn/dashboard-management/pkg/metadata"
)

// StreamName for area domain
var StreamName = fmt.Sprintf("%T", Area{})

type Area struct {
	areaID  uuid.UUID
	cityID  uuid.UUID
	name    string
	version int
	changes []*domain.Event
}

// New creates an Area
func New() Area {
	return Area{}
}

// FromHistory loads current aggregate root state by applying all events in order
func FromHistory(ctx context.Context, events []*domain.Event) (Area, error) {
	c := New()

	for _, domainEvent := range events {
		var e domain.RawEvent
		//........
		switch domainEvent.Type {
		case WasCreatedType:
			e = domainEvent.Payload.(*WasCreated)
		case WasRemovedType:
			e = domainEvent.Payload.(*WasRemoved)
		case WasUpdatedType:
			e = domainEvent.Payload.(*WasUpdated)

		default:
			return c, apperrors.Wrap(fmt.Errorf("unhandled client event %s", domainEvent.Type))
		}

		if err := c.transition(e); err != nil {
			return c, apperrors.Wrap(err)
		}

		c.version++
	}

	return c, nil
}

// ID returns aggregate root id
func (c Area) GetAreaID() uuid.UUID {
	return c.areaID
}

// Version returns current aggregate root version
func (c Area) Version() int {
	return c.version
}

// Changes returns all new applied events
func (c Area) Changes() []*domain.Event {
	return c.changes
}
func (c *Area) trackChange(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	var meta domain.EventMetadata
	if i, hasIdentity := identity.FromContext(ctx); hasIdentity {
		meta.Identity = i
	}
	if m, ok := metadata.FromContext(ctx); ok {
		meta.IPAddress = m.IPAddress
		meta.UserAgent = m.UserAgent
		meta.Referer = m.Referer
	}
	if !meta.IsEmpty() {
		event.WithMetadata(&meta)
	}

	c.changes = append(c.changes, event)
	c.version++

	return event, nil
}

// Create alters current area state and append changes to aggregate root
func (c *Area) Create(ctx context.Context, areaID uuid.UUID, name string, cityId uuid.UUID) error {
	e := &WasCreated{
		AreaID: areaID,
		Name:   name,
		CityID: cityId,
	}

	if err := c.transition(e); err != nil {
		return apperrors.Wrap(err)
	}

	domainEvent, err := domain.NewEventFromRawEvent(c.areaID, StreamName, c.version, e)
	if err != nil {
		return apperrors.Wrap(err)
	}

	if _, err := c.trackChange(ctx, domainEvent); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

// Remove alters current area state and append changes to aggregate root
func (c *Area) Remove(ctx context.Context) error {
	e := &WasRemoved{
		AreaID: c.areaID,
	}

	if err := c.transition(e); err != nil {
		return apperrors.Wrap(err)
	}

	domainEvent, err := domain.NewEventFromRawEvent(c.areaID, StreamName, c.version, e)
	if err != nil {
		return apperrors.Wrap(err)
	}

	if _, err := c.trackChange(ctx, domainEvent); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

// ChangeEmailAddress alters current user state and append changes to aggregate root
func (c *Area) Update(ctx context.Context, name string, cityId uuid.UUID) error {
	e := &WasUpdated{
		AreaID: c.areaID,
		Name:   name,
		CityID: cityId,
	}

	if err := c.transition(e); err != nil {
		return apperrors.Wrap(err)
	}

	domainEvent, err := domain.NewEventFromRawEvent(c.areaID, StreamName, c.version, e)
	if err != nil {
		return apperrors.Wrap(err)
	}

	if _, err := c.trackChange(ctx, domainEvent); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

func (c *Area) transition(e domain.RawEvent) error {
	switch e := e.(type) {

	case *WasCreated:
		c.areaID = e.AreaID
		c.cityID = e.CityID
		c.name = e.Name
	case *WasUpdated:
		if e.Name != "" {
			c.name = e.Name
		} else {
			e.Name = c.name
		}

		if e.CityID != uuid.Nil {
			c.cityID = e.CityID
		} else {
			e.CityID = c.cityID
		}

	}

	return nil
}
