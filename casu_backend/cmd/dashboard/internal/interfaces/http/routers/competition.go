package routers

import (
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/interfaces/http/handlers"
	"github.com/hsynrtn/dashboard-management/pkg/commandbus"
	"github.com/vardius/gorouter/v4"
)

func CompetitionRoutes(router gorouter.Router, competitionRepository persistence.CompetitionRepository, commandBus commandbus.CommandBus) gorouter.Router {
	router.GET("/competition", handlers.BuildListCompetitionsHandler(competitionRepository))
	router.GET("/competition/{id}", handlers.BuildGetCompetitionHandler(competitionRepository))
	router.POST("/dispatch/competition/{command}", handlers.BuildCompetitionCommandDispatchHandler(commandBus))
	return router
}
