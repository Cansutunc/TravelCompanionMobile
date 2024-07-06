package city

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hsynrtn/dashboard-management/pkg/domain"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/identity"
	"github.com/hsynrtn/dashboard-management/pkg/metadata"
)

// StreamName for city domain
var StreamName = fmt.Sprintf("%T", City{})

type City struct {
	cityID          uuid.UUID
	title           string
	createdUserId   string
	createdUsername string
	description     string
	maxPerson       int
	allPerson       []string
	version         int
	changes         []*domain.Event
}

// New creates an City
func New() City {
	return City{}
}

// FromHistory loads current aggregate root state by applying all events in order
func FromHistory(ctx context.Context, events []*domain.Event) (City, error) {
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
func (c City) GetCityID() uuid.UUID {
	return c.cityID
}

// Version returns current aggregate root version
func (c City) Version() int {
	return c.version
}

// Changes returns all new applied events
func (c City) Changes() []*domain.Event {
	return c.changes
}
func (c *City) trackChange(ctx context.Context, event *domain.Event) (*domain.Event, error) {
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

// Create alters current city state and append changes to aggregate root
func (c *City) Create(ctx context.Context, cityID uuid.UUID, title string, description string, maxPerson int, createdUserId string, createdUsername string) error {
	e := &WasCreated{
		CityID:          cityID,
		Title:           title,
		Description:     description,
		MaxPerson:       maxPerson,
		CreatedUsername: createdUsername,
		CreatedUserId:   createdUserId,
	}

	if err := c.transition(e); err != nil {
		return apperrors.Wrap(err)
	}

	domainEvent, err := domain.NewEventFromRawEvent(c.cityID, StreamName, c.version, e)
	if err != nil {
		return apperrors.Wrap(err)
	}

	if _, err := c.trackChange(ctx, domainEvent); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

// Remove alters current city state and append changes to aggregate root
func (c *City) Remove(ctx context.Context) error {
	e := &WasRemoved{
		CityID: c.cityID,
	}

	if err := c.transition(e); err != nil {
		return apperrors.Wrap(err)
	}

	domainEvent, err := domain.NewEventFromRawEvent(c.cityID, StreamName, c.version, e)
	if err != nil {
		return apperrors.Wrap(err)
	}

	if _, err := c.trackChange(ctx, domainEvent); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

// ChangeEmailAddress alters current user state and append changes to aggregate root
func (c *City) Update(ctx context.Context, name string, code string, countryId uuid.UUID) error {
	e := &WasUpdated{
		CityID:    c.cityID,
		Name:      name,
		Code:      code,
		CountryID: countryId,
	}

	if err := c.transition(e); err != nil {
		return apperrors.Wrap(err)
	}

	domainEvent, err := domain.NewEventFromRawEvent(c.cityID, StreamName, c.version, e)
	if err != nil {
		return apperrors.Wrap(err)
	}

	if _, err := c.trackChange(ctx, domainEvent); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

func (c *City) transition(e domain.RawEvent) error {
	switch e := e.(type) {

	case *WasCreated:
		c.cityID = e.CityID
		c.title = e.Title
		c.description = e.Description
		c.maxPerson = e.MaxPerson
		c.createdUsername = e.CreatedUsername
		c.createdUserId = e.CreatedUserId

	}

	return nil
}
