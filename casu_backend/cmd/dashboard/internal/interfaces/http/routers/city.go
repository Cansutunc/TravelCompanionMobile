package routers

import (
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/interfaces/http/handlers"
	"github.com/hsynrtn/dashboard-management/pkg/commandbus"
	"github.com/vardius/gorouter/v4"
)

func CityRoutes(router gorouter.Router, cityRepository persistence.CityRepository, commandBus commandbus.CommandBus) gorouter.Router {
	router.GET("/city", handlers.BuildListCitiesHandler(cityRepository))
	router.GET("/city/{id}", handlers.BuildGetCityHandler(cityRepository))
	router.GET("/city/accept", handlers.BuildAcceptHandler(cityRepository))
	router.GET("/city/acceptAll", handlers.BuildGetAcceptedHandler(cityRepository))
	router.GET("/city/remove", handlers.BuildRemoveHandler(cityRepository))
	router.POST("/dispatch/city/{command}", handlers.BuildCityCommandDispatchHandler(commandBus))
	return router
}
