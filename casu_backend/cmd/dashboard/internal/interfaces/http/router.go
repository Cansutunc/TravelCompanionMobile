package http

import (
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/application/config"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/application/services"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/interfaces/http/handlers"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/interfaces/http/routers"
	"net/http"
	"time"

	httpmiddleware "github.com/hsynrtn/dashboard-management/pkg/http/middleware"
	"github.com/hsynrtn/dashboard-management/pkg/http/response/json"

	"github.com/vardius/gorouter/v4"
)

// NewRouter provides new router
func NewRouter(
	cfg *config.Config,
	container *services.ServiceContainer,

) http.Handler {

	// Global middleware
	router := gorouter.New(
		httpmiddleware.Recover(),
		httpmiddleware.WithMetadata(),
		httpmiddleware.Logger(),
		httpmiddleware.XSS(),
		httpmiddleware.HSTS(),

		httpmiddleware.CORS(
			cfg.HTTP.Origins,
			cfg.App.Environment == "development",
		),
		httpmiddleware.LimitRequestBody(int64(10<<20)), // 10 MB is a lot of text.
		httpmiddleware.Metrics(),
		httpmiddleware.RateLimit(10, 10, 3*time.Minute), // 5 of requests per second with bursts of at most 10 requests
	)
	router.NotFound(json.NotFound())
	router.NotAllowed(json.NotAllowed())

	router = routers.CountryRoutes(router, container.CountryPersistenceRepository, container.CommandBus)
	router = routers.CityRoutes(router, container.CityPersistenceRepository, container.CommandBus)
	router = routers.AreaRoutes(router, container.AreaPersistenceRepository, container.CommandBus)
	router = routers.CompetitionRoutes(router, container.CompetitionPersistenceRepository, container.CommandBus)
	router = routers.MatchRoutes(router, container.MatchPersistenceRepository, container.CommandBus)
	router = routers.BookRoutes(router, container.BookPersistenceRepository, container.CommandBus)
	router = routers.GeneralRoutes(router)

	mainRouter := gorouter.New()
	mainRouter.NotFound(json.NotFound())
	mainRouter.NotAllowed(json.NotAllowed())

	// We do not want to apply middleware for this handlers
	// Liveness probes are to indicate that your application is running
	mainRouter.GET("/health", handlers.BuildLivenessHandler())
	// Readiness is meant to check if your application is ready to serve traffic
	mainRouter.GET("/readiness", handlers.BuildReadinessHandler(container.Mongo))

	mainRouter.Mount("/v1", router)

	return mainRouter
}
