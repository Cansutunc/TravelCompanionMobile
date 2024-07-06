package area

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
	CreateArea = "area-create"
	RemoveArea = "area-remove"
	UpdateArea = "area-update"
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
	case CreateArea:
		var command Create
		if err := json.Unmarshal(payload, &command); err != nil {
			return command, apperrors.Wrap(err)
		}
		return command, nil
	case RemoveArea:
		var command Remove
		if err := json.Unmarshal(payload, &command); err != nil {
			return command, apperrors.Wrap(err)
		}
		return command, nil
	case UpdateArea:
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
	Name   string    `json:"name"`
	CityID uuid.UUID `json:"city_id"`
}

// Remove command
type Remove struct {
	ID uuid.UUID `json:"id"`
}
type Update struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name,omitempty"`
	CityID uuid.UUID `json:"city_id,omitempty"`
}

// OnCreate creates command handler
func OnCreate(repository Repository, areaRepository persistence.AreaRepository) commandbus.CommandHandler {
	fn := func(ctx context.Context, command domain.Command) error {
		c, ok := command.(Create)
		if !ok {
			return apperrors.New("invalid command")
		}

		_, err := areaRepository.GetByName(ctx, c.Name)
		if err != nil && !errors.Is(err, apperrors.ErrNotFound) {
			return apperrors.Wrap(err)
		}
		var u Area
		if err == nil {
			return apperrors.Wrap(apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrAlreadyExist)))
		} else {
			id, err := uuid.NewRandom()
			if err != nil {
				return apperrors.Wrap(err)
			}
			u = New()
			if err := u.Create(ctx, id, c.Name, c.CityID); err != nil {
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

		area, err := repository.Get(ctx, c.ID)
		if err != nil {
			return apperrors.Wrap(err)
		}

		if err := area.Remove(ctx); err != nil {
			return apperrors.Wrap(fmt.Errorf("%w: Error when removing token: %s", apperrors.ErrInternal, err))
		}

		if err := repository.Save(executioncontext.WithFlag(ctx, executioncontext.LIVE), area); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return fn
}

// OnUpdate creates command handler
func OnUpdate(repository Repository, areaRepository persistence.AreaRepository) commandbus.CommandHandler {
	fn := func(ctx context.Context, command domain.Command) error {
		c, ok := command.(Update)
		if !ok {
			return apperrors.New("invalid command")
		}

		u, err := repository.Get(ctx, c.ID)
		if err != nil {
			return apperrors.Wrap(err)
		}

		if err := u.Update(ctx, c.Name, c.CityID); err != nil {
			return apperrors.Wrap(err)
		}

		if err := repository.Save(executioncontext.WithFlag(ctx, executioncontext.LIVE), u); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return fn
}
