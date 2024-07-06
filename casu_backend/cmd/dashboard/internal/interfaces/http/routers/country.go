package routers

import (
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/interfaces/http/handlers"
	"github.com/hsynrtn/dashboard-management/pkg/commandbus"
	"github.com/vardius/gorouter/v4"
)

func CountryRoutes(router gorouter.Router, countryRepository persistence.CountryRepository, commandBus commandbus.CommandBus) gorouter.Router {
	router.GET("/country", handlers.BuildListCountriesHandler(countryRepository))
	router.GET("/country/{id}", handlers.BuildGetCountryHandler(countryRepository))
	router.POST("/dispatch/country/{command}", handlers.BuildCountryCommandDispatchHandler(commandBus))
	return router
}
