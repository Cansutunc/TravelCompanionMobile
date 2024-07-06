package eventhandler

import (
	"context"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/book"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence"
	"github.com/hsynrtn/dashboard-management/pkg/domain"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/eventbus"
	"time"
)

// WhenBookWasRemoved handles event
func WhenBookWasRemoved(bookRepository persistence.BookRepository) eventbus.EventHandler {
	fn := func(parentCtx context.Context, event *domain.Event) error {
		ctx, cancel := context.WithTimeout(parentCtx, time.Second*120)
		defer cancel()

		e := event.Payload.(*book.WasRemoved)

		if err := bookRepository.Delete(ctx, e.BookID.String()); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return fn
}
