package competition

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hsynrtn/dashboard-management/pkg/domain"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/identity"
	"github.com/hsynrtn/dashboard-management/pkg/metadata"
)

// StreamName for competition domain
var StreamName = fmt.Sprintf("%T", Competition{})

type Competition struct {
	competitionID uuid.UUID
	countryID     uuid.UUID
	name          string
	code          string
	version       int
	changes       []*domain.Event
}

// New creates an Competition
func New() Competition {
	return Competition{}
}

// FromHistory loads current aggregate root state by applying all events in order
func FromHistory(ctx context.Context, events []*domain.Event) (Competition, error) {
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
func (c Competition) GetCompetitionID() uuid.UUID {
	return c.competitionID
}

// Version returns current aggregate root version
func (c Competition) Version() int {
	return c.version
}

// Changes returns all new applied events
func (c Competition) Changes() []*domain.Event {
	return c.changes
}
func (c *Competition) trackChange(ctx context.Context, event *domain.Event) (*domain.Event, error) {
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

// Create alters current competition state and append changes to aggregate root
func (c *Competition) Create(ctx context.Context, competitionID uuid.UUID, name string, sporType string, region string, competitionType string) error {
	e := &WasCreated{
		CompetitionID:   competitionID,
		Name:            name,
		SporType:        sporType,
		Region:          region,
		CompetitionType: competitionType,
	}

	if err := c.transition(e); err != nil {
		return apperrors.Wrap(err)
	}

	domainEvent, err := domain.NewEventFromRawEvent(c.competitionID, StreamName, c.version, e)
	if err != nil {
		return apperrors.Wrap(err)
	}

	if _, err := c.trackChange(ctx, domainEvent); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

// Remove alters current competition state and append changes to aggregate root
func (c *Competition) Remove(ctx context.Context) error {
	e := &WasRemoved{
		CompetitionID: c.competitionID,
	}

	if err := c.transition(e); err != nil {
		return apperrors.Wrap(err)
	}

	domainEvent, err := domain.NewEventFromRawEvent(c.competitionID, StreamName, c.version, e)
	if err != nil {
		return apperrors.Wrap(err)
	}

	if _, err := c.trackChange(ctx, domainEvent); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

// ChangeEmailAddress alters current user state and append changes to aggregate root
func (c *Competition) Update(ctx context.Context, name string, code string, countryId uuid.UUID) error {
	e := &WasUpdated{
		CompetitionID: c.competitionID,
		Name:          name,
		Code:          code,
		CountryID:     countryId,
	}

	if err := c.transition(e); err != nil {
		return apperrors.Wrap(err)
	}

	domainEvent, err := domain.NewEventFromRawEvent(c.competitionID, StreamName, c.version, e)
	if err != nil {
		return apperrors.Wrap(err)
	}

	if _, err := c.trackChange(ctx, domainEvent); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

func (c *Competition) transition(e domain.RawEvent) error {
	switch e := e.(type) {

	case *WasCreated:
		c.competitionID = e.CompetitionID
	case *WasUpdated:
		if e.Name != "" {
			c.name = e.Name
		} else {
			e.Name = c.name
		}

		if e.Code != "" {
			c.code = e.Code
		} else {
			e.Code = c.code
		}

		if e.CountryID != uuid.Nil {
			c.countryID = e.CountryID
		} else {
			e.CountryID = c.countryID
		}

	}

	return nil
}
