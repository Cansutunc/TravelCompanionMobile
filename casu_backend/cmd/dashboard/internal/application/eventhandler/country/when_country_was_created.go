package eventhandler

import (
	"context"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/country"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence"
	"github.com/hsynrtn/dashboard-management/pkg/domain"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/eventbus"
	"time"
)

// WhenCountryWasCreated handles event
func WhenCountryWasCreated(repository persistence.CountryRepository) eventbus.EventHandler {
	fn := func(parentCtx context.Context, event *domain.Event) error {
		ctx, cancel := context.WithTimeout(parentCtx, time.Second*120)
		defer cancel()

		e := event.Payload.(*country.WasCreated)

		if err := repository.Add(ctx, e); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return fn
}