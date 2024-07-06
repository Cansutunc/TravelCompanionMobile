package eventhandler

import (
	"context"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/area"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence"
	"github.com/hsynrtn/dashboard-management/pkg/domain"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/eventbus"
	"time"
)

// WhenCityWasRemoved handles event
func WhenAreaWasRemoved(cityRepository persistence.AreaRepository) eventbus.EventHandler {
	fn := func(parentCtx context.Context, event *domain.Event) error {
		ctx, cancel := context.WithTimeout(parentCtx, time.Second*120)
		defer cancel()

		e := event.Payload.(*area.WasRemoved)

		if err := cityRepository.Delete(ctx, e.AreaID.String()); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return fn
}
