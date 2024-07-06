package country

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hsynrtn/dashboard-management/pkg/domain"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/identity"
	"github.com/hsynrtn/dashboard-management/pkg/metadata"
)

// StreamName for country domain
var StreamName = fmt.Sprintf("%T", Country{})

type Country struct {
	countryID uuid.UUID
	version   int
	changes   []*domain.Event
	userId    string
	country   string
	province  string
	username  string
}

// New creates an Country
func New() Country {
	return Country{}
}

// Create alters current user state and append changes to aggregate root
func (c *Country) Create(
	ctx context.Context,
	id uuid.UUID,
	userId string,
	province string,
	country string,
	username string,

) error {
	e := &WasCreated{
		CountryID: id,
		Province:  province,
		UserId:    userId,
		Country:   country,
		UserName:  username,
	}

	if err := c.transition(e); err != nil {
		return apperrors.Wrap(err)
	}

	domainEvent, err := domain.NewEventFromRawEvent(c.countryID, StreamName, c.version, e)
	if err != nil {
		return apperrors.Wrap(err)
	}

	if _, err := c.trackChange(ctx, domainEvent); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

// Remove alters current token state and append changes to aggregate root
func (c *Country) Remove(ctx context.Context) error {
	e := &WasRemoved{
		CountryID: c.countryID,
	}

	if err := c.transition(e); err != nil {
		return apperrors.Wrap(err)
	}

	domainEvent, err := domain.NewEventFromRawEvent(c.countryID, StreamName, c.version, e)
	if err != nil {
		return apperrors.Wrap(err)
	}

	if _, err := c.trackChange(ctx, domainEvent); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

// ChangeEmailAddress alters current user state and append changes to aggregate root
func (c *Country) Update(ctx context.Context, userId string, province string, country string, userName string) error {
	e := &WasUpdated{
		CountryID: c.countryID,
		UserId:    userId,
		Province:  province,
		Country:   country,
		UserName:  userName,
	}

	if err := c.transition(e); err != nil {
		return apperrors.Wrap(err)
	}

	domainEvent, err := domain.NewEventFromRawEvent(c.countryID, StreamName, c.version, e)
	if err != nil {
		return apperrors.Wrap(err)
	}

	if _, err := c.trackChange(ctx, domainEvent); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

// FromHistory loads current aggregate root state by applying all events in order
func FromHistory(ctx context.Context, events []*domain.Event) (Country, error) {
	c := New()

	for _, domainEvent := range events {
		var e domain.RawEvent

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
func (c Country) GetCountryID() uuid.UUID {
	return c.countryID
}

// Version returns current aggregate root version
func (c Country) Version() int {
	return c.version
}

// Changes returns all new applied events
func (c Country) Changes() []*domain.Event {
	return c.changes
}

func (c *Country) trackChange(ctx context.Context, event *domain.Event) (*domain.Event, error) {
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

func (c *Country) transition(e domain.RawEvent) error {
	switch e := e.(type) {
	case *WasUpdated:
		if e.UserId != "" {
			c.userId = e.UserId
		} else {
			e.UserId = c.userId
		}
		if e.Country != "" {
			c.country = e.Country
		} else {
			e.Country = c.country
		}

	case *WasCreated:
		c.countryID = e.CountryID
		c.userId = e.UserId
		c.country = e.Country
		c.province = e.Province
		c.username = e.UserName

	}

	return nil
}
