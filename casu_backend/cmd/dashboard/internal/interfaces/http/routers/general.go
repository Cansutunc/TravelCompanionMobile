package routers

import (
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/interfaces/http/handlers"
	"github.com/vardius/gorouter/v4"
)

func GeneralRoutes(router gorouter.Router) gorouter.Router {
	router.GET("/get-sport-types", handlers.BuildListSportTypes())
	router.GET("/get-blood-types", handlers.BuildListBloodtTypes())

	return router
}
