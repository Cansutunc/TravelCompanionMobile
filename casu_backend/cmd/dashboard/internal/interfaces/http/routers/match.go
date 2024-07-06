package routers

import (
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/interfaces/http/handlers"
	"github.com/hsynrtn/dashboard-management/pkg/commandbus"
	"github.com/vardius/gorouter/v4"
)

func MatchRoutes(router gorouter.Router, matchRepository persistence.MatchRepository, commandBus commandbus.CommandBus) gorouter.Router {
	router.GET("/match", handlers.BuildListMatchesHandler(matchRepository))
	router.GET("/match/{id}", handlers.BuildGetMatchHandler(matchRepository))
	router.POST("/dispatch/match/{command}", handlers.BuildMatchCommandDispatchHandler(commandBus))
	return router
}
