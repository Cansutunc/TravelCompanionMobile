package eventhandler

import (
	"context"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/match"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence"
	"github.com/hsynrtn/dashboard-management/pkg/domain"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/eventbus"
	"time"
)

// WhenMatchWasUpdated handles event
func WhenMatchWasUpdated(matchRepository persistence.MatchRepository) eventbus.EventHandler {
	fn := func(parentCtx context.Context, event *domain.Event) error {
		ctx, cancel := context.WithTimeout(parentCtx, time.Second*120)
		defer cancel()

		e := event.Payload.(*match.WasUpdated)

		if err := matchRepository.Update(ctx, e.MatchID.String(), e.Name, e.Code, e.CountryID.String()); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return fn
}
