package routers

import (
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/interfaces/http/handlers"
	"github.com/hsynrtn/dashboard-management/pkg/commandbus"
	"github.com/vardius/gorouter/v4"
)

func AreaRoutes(router gorouter.Router, areaRepository persistence.AreaRepository, commandBus commandbus.CommandBus) gorouter.Router {
	router.GET("/area", handlers.BuildListAreaHandler(areaRepository))
	router.GET("/area/{id}", handlers.BuildGetAreaHandler(areaRepository))
	router.POST("/dispatch/area/{command}", handlers.BuildAreaCommandDispatchHandler(commandBus))
	return router
}
