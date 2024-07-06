package competition

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence"
	"github.com/hsynrtn/dashboard-management/pkg/commandbus"
	"github.com/hsynrtn/dashboard-management/pkg/domain"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/executioncontext"
)

const (
	CreateCompetition = "competition-create"
	RemoveCompetition = "competition-remove"
	UpdateCompetition = "competition-update"
)

var (
	CreateName = (Create{}).GetName()
	RemoveName = (Remove{}).GetName()
	UpdateName = (Update{}).GetName()
)

// GetName returns command name
func (c Create) GetName() string {
	return fmt.Sprintf("%T", c)
}
func (c Remove) GetName() string {
	return fmt.Sprintf("%T", c)
}
func (c Update) GetName() string {
	return fmt.Sprintf("%T", c)
}

func NewCommandFromPayload(contract string, payload []byte) (domain.Command, error) {
	switch contract {
	case CreateCompetition:
		var command Create
		if err := json.Unmarshal(payload, &command); err != nil {
			return command, apperrors.Wrap(err)
		}
		return command, nil
	case RemoveCompetition:
		var command Remove
		if err := json.Unmarshal(payload, &command); err != nil {
			return command, apperrors.Wrap(err)
		}
		return command, nil
	case UpdateCompetition:
		var command Update
		if err := json.Unmarshal(payload, &command); err != nil {
			return command, apperrors.Wrap(err)
		}
		return command, nil
	default:
		return nil, apperrors.Wrap(fmt.Errorf("invalid command contract: %s", contract))
	}

}

// Create command
type Create struct {
	Name            string `json:"name"`
	SporType        string `json:"spor_type"`
	Region          string `json:"region"`
	CompetitionType string `json:"competition_type"`
}

// Remove command
type Remove struct {
	ID uuid.UUID `json:"id"`
}
type Update struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name,omitempty"`
	Code      string    `json:"code,omitempty"`
	CountryID uuid.UUID `json:"country_id,omitempty"`
}

// OnCreate creates command handler
func OnCreate(repository Repository, competitionRepository persistence.CompetitionRepository) commandbus.CommandHandler {
	fn := func(ctx context.Context, command domain.Command) error {
		c, ok := command.(Create)
		if !ok {
			return apperrors.New("invalid command")
		}

		_, err := competitionRepository.GetByName(ctx, c.Name)
		if err != nil && !errors.Is(err, apperrors.ErrNotFound) {
			return apperrors.Wrap(err)
		}
		var u Competition
		if err == nil {
			return apperrors.Wrap(apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrAlreadyExist)))
		} else {
			id, err := uuid.NewRandom()
			if err != nil {
				return apperrors.Wrap(err)
			}
			u = New()
			if err := u.Create(ctx, id, c.Name, c.SporType, c.Region, c.CompetitionType); err != nil {
				return apperrors.Wrap(err)
			}
		}

		if err := repository.Save(executioncontext.WithFlag(ctx, executioncontext.LIVE), u); err != nil {
			return apperrors.Wrap(err)
		}
		return nil
	}

	return fn
}

// OnRemove creates command handler
func OnRemove(repository Repository) commandbus.CommandHandler {
	fn := func(ctx context.Context, command domain.Command) error {
		c, ok := command.(Remove)
		if !ok {
			return apperrors.New("invalid command")
		}

		competition, err := repository.Get(ctx, c.ID)
		if err != nil {
			return apperrors.Wrap(err)
		}

		if err := competition.Remove(ctx); err != nil {
			return apperrors.Wrap(fmt.Errorf("%w: Error when removing token: %s", apperrors.ErrInternal, err))
		}

		if err := repository.Save(executioncontext.WithFlag(ctx, executioncontext.LIVE), competition); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return fn
}

// OnUpdate creates command handler
func OnUpdate(repository Repository, competitionRepository persistence.CompetitionRepository) commandbus.CommandHandler {
	fn := func(ctx context.Context, command domain.Command) error {
		c, ok := command.(Update)
		if !ok {
			return apperrors.New("invalid command")
		}

		u, err := repository.Get(ctx, c.ID)
		if err != nil {
			return apperrors.Wrap(err)
		}

		if err := u.Update(ctx, c.Name, c.Code, c.CountryID); err != nil {
			return apperrors.Wrap(err)
		}

		if err := repository.Save(executioncontext.WithFlag(ctx, executioncontext.LIVE), u); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return fn
}
