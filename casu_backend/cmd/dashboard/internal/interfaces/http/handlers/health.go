package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"

	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	httpjson "github.com/hsynrtn/dashboard-management/pkg/http/response/json"
)

// BuildLivenessHandler provides liveness handler
func BuildLivenessHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}

	return http.HandlerFunc(fn)
}

// BuildReadinessHandler provides readiness handler
func BuildReadinessHandler(mongoConn *mongo.Client) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if mongoConn != nil {
			if err := mongoConn.Ping(r.Context(), nil); err != nil {
				return apperrors.Wrap(err)
			}
		}
		w.WriteHeader(http.StatusNoContent)

		return nil
	}

	return httpjson.HandlerFunc(fn)
}
