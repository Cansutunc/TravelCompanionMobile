package eventhandler

import (
	"context"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/competition"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence"
	"github.com/hsynrtn/dashboard-management/pkg/domain"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/eventbus"
	"time"
)

// WhenCompetitionWasRemoved handles event
func WhenCompetitionWasRemoved(competitionRepository persistence.CompetitionRepository) eventbus.EventHandler {
	fn := func(parentCtx context.Context, event *domain.Event) error {
		ctx, cancel := context.WithTimeout(parentCtx, time.Second*120)
		defer cancel()

		e := event.Payload.(*competition.WasRemoved)

		if err := competitionRepository.Delete(ctx, e.CompetitionID.String()); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return fn
}
