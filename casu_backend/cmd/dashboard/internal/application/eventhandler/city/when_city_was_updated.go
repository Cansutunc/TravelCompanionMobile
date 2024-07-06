package eventhandler

import (
	"context"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/city"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence"
	"github.com/hsynrtn/dashboard-management/pkg/domain"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/eventbus"
	"time"
)

// WhenCityWasUpdated handles event
func WhenCityWasUpdated(cityRepository persistence.CityRepository) eventbus.EventHandler {
	fn := func(parentCtx context.Context, event *domain.Event) error {
		ctx, cancel := context.WithTimeout(parentCtx, time.Second*120)
		defer cancel()

		e := event.Payload.(*city.WasUpdated)

		if err := cityRepository.Update(ctx, e.CityID.String(), e.Name, e.Code, e.CountryID.String()); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return fn
}
