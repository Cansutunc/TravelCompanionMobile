package errors

import (
	"context"
	"errors"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"net/http"

	mtd "github.com/hsynrtn/dashboard-management/pkg/metadata"
)

type HttpError struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Error     string `json:"error"`
	RequestID string `json:"request_id,omitempty"`
}

func NewHttpError(ctx context.Context, err error) *HttpError {
	code := http.StatusInternalServerError

	switch {
	case errors.Is(err, apperrors.ErrInvalid):
		code = http.StatusBadRequest
	case errors.Is(err, apperrors.ErrUnauthorized):
		code = http.StatusUnauthorized
	case errors.Is(err, apperrors.ErrForbidden):
		code = http.StatusForbidden
	case errors.Is(err, apperrors.ErrNotFound):
		code = http.StatusNotFound
	case errors.Is(err, apperrors.ErrAlreadyExist):
		code = http.StatusBadRequest

	case errors.Is(err, apperrors.ErrTimeout):
		code = http.StatusRequestTimeout
	case errors.Is(err, apperrors.ErrTemporaryDisabled):
		code = http.StatusServiceUnavailable
	case errors.Is(err, apperrors.ErrInternal):
		code = http.StatusInternalServerError
	}

	httpError := &HttpError{
		Code:    code,
		Message: http.StatusText(code),
		Error:   err.Error(),
	}

	if m, ok := mtd.FromContext(ctx); ok {
		httpError.RequestID = m.TraceID
		m.Err = err
	}

	return httpError
}
