package country

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
	CreateCountry = "country-create"
	RemoveCountry = "country-remove"
	UpdateCountry = "country-update"
)

var (
	CreateName = (Create{}).GetName()
	RemoveName = (Remove{}).GetName()
	UpdateName = (Update{}).GetName()
)

func NewCommandFromPayload(contract string, payload []byte) (domain.Command, error) {
	switch contract {
	case CreateCountry:
		var command Create
		if err := json.Unmarshal(payload, &command); err != nil {
			return command, apperrors.Wrap(err)
		}
		return command, nil
	case RemoveCountry:
		var command Remove
		if err := json.Unmarshal(payload, &command); err != nil {
			return command, apperrors.Wrap(err)
		}
		return command, nil

	case UpdateCountry:
		var command Update
		if err := json.Unmarshal(payload, &command); err != nil {
			return command, apperrors.Wrap(err)
		}
		return command, nil

	default:
		return nil, apperrors.Wrap(fmt.Errorf("invalid command contract: %s", contract))
	}

}

// Create command, creates access token for user
type Create struct {
	Province string `json:"province"`
	Country  string `json:"country"`
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}

// Remove command
type Remove struct {
	ID uuid.UUID `json:"id"`
}

// Update command
type Update struct {
	ID       uuid.UUID `json:"id"`
	Province string    `json:"province,omitempty"`
	Country  string    `json:"country,omitempty"`
	UserId   string    `json:"user_id,omitempty"`
	UserName string    `json:"user_name,omitempty"`
}

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

// OnCreate creates command handler
func OnCreate(repository Repository, countryRepository persistence.CountryRepository) commandbus.CommandHandler {
	fn := func(ctx context.Context, command domain.Command) error {
		c, ok := command.(Create)
		if !ok {
			return apperrors.New("invalid command")
		}

		_, err := countryRepository.GetByUserIdAndProvince(ctx, c.UserId, c.Province)
		if err != nil && !errors.Is(err, apperrors.ErrNotFound) {
			return apperrors.Wrap(err)
		}
		var u Country
		if err == nil {
			return apperrors.Wrap(apperrors.ErrAlreadyExist)
		} else {
			id, err := uuid.NewRandom()
			if err != nil {
				return apperrors.Wrap(err)
			}
			u = New()
			if err := u.Create(ctx, id, c.UserId, c.Province, c.Country, c.UserName); err != nil {
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

		country, err := repository.Get(ctx, c.ID)
		if err != nil {
			return apperrors.Wrap(err)
		}

		if err := country.Remove(ctx); err != nil {
			return apperrors.Wrap(fmt.Errorf("%w: Error when removing token: %s", apperrors.ErrInternal, err))
		}

		if err := repository.Save(executioncontext.WithFlag(ctx, executioncontext.LIVE), country); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return fn
}

// OnUpdate creates command handler
func OnUpdate(repository Repository, countryRepository persistence.CountryRepository) commandbus.CommandHandler {
	fn := func(ctx context.Context, command domain.Command) error {
		c, ok := command.(Update)
		if !ok {
			return apperrors.New("invalid command")
		}

		if c.Province != "" {
			country, err := countryRepository.GetByUserIdAndProvince(ctx, c.UserId, c.Province)
			if err != nil && !errors.Is(err, apperrors.ErrNotFound) {
				return apperrors.Wrap(err)
			}
			if err == nil {
				if country.GetCountryID() != c.ID.String() {
					return apperrors.Wrap(apperrors.ErrAlreadyExist)
				}

			}
		}

		u, err := repository.Get(ctx, c.ID)
		if err != nil {
			return apperrors.Wrap(err)
		}

		if err := u.Update(ctx, c.UserId, c.Country, c.Province, c.UserName); err != nil {
			return apperrors.Wrap(err)
		}

		if err := repository.Save(executioncontext.WithFlag(ctx, executioncontext.LIVE), u); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return fn
}
