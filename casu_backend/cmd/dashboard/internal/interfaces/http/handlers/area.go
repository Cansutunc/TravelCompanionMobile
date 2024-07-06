package handlers

import (
	"fmt"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/area"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence"
	"github.com/hsynrtn/dashboard-management/pkg/commandbus"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	httpjson "github.com/hsynrtn/dashboard-management/pkg/http/response/json"
	"github.com/vardius/gorouter/v4/context"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
)

// BuildListAreaHandler lists client credentials by user ID
func BuildListAreaHandler(repository persistence.AreaRepository) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		pageInt, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)
		limitInt, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
		page := int64(math.Max(float64(pageInt), 1))
		limit := int64(math.Max(float64(limitInt), 5))

		totalUsers, err := repository.Count(r.Context())
		if err != nil {
			return apperrors.Wrap(err)
		}

		offset := (page * limit) - limit

		paginatedList := struct {
			Areas []persistence.Area `json:"areas"`
			Page  int64              `json:"page"`
			Limit int64              `json:"limit"`
			Total int64              `json:"total"`
		}{
			Page:  page,
			Limit: limit,
			Total: totalUsers,
		}

		if totalUsers < 1 || offset > (totalUsers-1) {
			if err := httpjson.JSON(r.Context(), w, http.StatusOK, paginatedList); err != nil {
				return apperrors.Wrap(err)
			}
			return nil
		}

		paginatedList.Areas, err = repository.FindAll(r.Context(), limit, offset)
		if err != nil {
			return apperrors.Wrap(err)
		}

		if err := httpjson.JSON(r.Context(), w, http.StatusOK, paginatedList); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return httpjson.HandlerFunc(fn)
}

// BuildGetAreaHandler
func BuildGetAreaHandler(repository persistence.AreaRepository) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		params, ok := context.Parameters(r.Context())
		if !ok {
			return apperrors.Wrap(ErrInvalidURLParams)
		}

		u, err := repository.Get(r.Context(), params.Value("id"))
		if err != nil {
			return apperrors.Wrap(err)
		}

		if err := httpjson.JSON(r.Context(), w, http.StatusOK, u); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return httpjson.HandlerFunc(fn)
}

// BuildAreaCommandDispatchHandler
func BuildAreaCommandDispatchHandler(cb commandbus.CommandBus) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if r.Body == nil {
			return fmt.Errorf("%w: %v", apperrors.ErrInvalid, ErrEmptyRequestBody)
		}

		params, ok := context.Parameters(r.Context())
		if !ok {
			return fmt.Errorf("%w: %v", apperrors.ErrInvalid, ErrInvalidURLParams)
		}

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return apperrors.Wrap(err)
		}

		c, err := area.NewCommandFromPayload(params.Value("command"), body)
		if err != nil {
			return apperrors.Wrap(err)
		}

		if err := cb.Publish(r.Context(), c); err != nil {
			return apperrors.Wrap(err)
		}

		if err := httpjson.JSON(r.Context(), w, http.StatusCreated, nil); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return httpjson.HandlerFunc(fn)
}
