package book

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hsynrtn/dashboard-management/pkg/domain"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/identity"
	"github.com/hsynrtn/dashboard-management/pkg/metadata"
)

// StreamName for book domain
var StreamName = fmt.Sprintf("%T", Book{})

type Book struct {
	bookID  uuid.UUID
	version int
	changes []*domain.Event
	name    string
	code    string
}

// New creates an Book
func New() Book {
	return Book{}
}

// Create alters current user state and append changes to aggregate root
func (c *Book) Create(
	ctx context.Context,
	id uuid.UUID,
	name string,
	code string,
	author string,
	homepage string,

) error {
	e := &WasCreated{
		BookID:   id,
		Name:     name,
		Code:     code,
		Author:   author,
		Homepage: homepage,
	}

	if err := c.transition(e); err != nil {
		return apperrors.Wrap(err)
	}

	domainEvent, err := domain.NewEventFromRawEvent(c.bookID, StreamName, c.version, e)
	if err != nil {
		return apperrors.Wrap(err)
	}

	if _, err := c.trackChange(ctx, domainEvent); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

// Remove alters current token state and append changes to aggregate root
func (c *Book) Remove(ctx context.Context) error {
	e := &WasRemoved{
		BookID: c.bookID,
	}

	if err := c.transition(e); err != nil {
		return apperrors.Wrap(err)
	}

	domainEvent, err := domain.NewEventFromRawEvent(c.bookID, StreamName, c.version, e)
	if err != nil {
		return apperrors.Wrap(err)
	}

	if _, err := c.trackChange(ctx, domainEvent); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

// ChangeEmailAddress alters current user state and append changes to aggregate root
func (c *Book) Update(ctx context.Context, name string, code string) error {
	e := &WasUpdated{
		BookID: c.bookID,
		Name:   name,
		Code:   code,
	}

	if err := c.transition(e); err != nil {
		return apperrors.Wrap(err)
	}

	domainEvent, err := domain.NewEventFromRawEvent(c.bookID, StreamName, c.version, e)
	if err != nil {
		return apperrors.Wrap(err)
	}

	if _, err := c.trackChange(ctx, domainEvent); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

// FromHistory loads current aggregate root state by applying all events in order
func FromHistory(ctx context.Context, events []*domain.Event) (Book, error) {
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
func (c Book) GetBookID() uuid.UUID {
	return c.bookID
}

// Version returns current aggregate root version
func (c Book) Version() int {
	return c.version
}

// Changes returns all new applied events
func (c Book) Changes() []*domain.Event {
	return c.changes
}

func (c *Book) trackChange(ctx context.Context, event *domain.Event) (*domain.Event, error) {
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

func (c *Book) transition(e domain.RawEvent) error {
	switch e := e.(type) {
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

	case *WasCreated:
		c.bookID = e.BookID
		c.name = e.Name
		c.code = e.Code

	}

	return nil
}
