package routers

import (
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/interfaces/http/handlers"
	"github.com/hsynrtn/dashboard-management/pkg/commandbus"
	"github.com/vardius/gorouter/v4"
)

func BookRoutes(router gorouter.Router, bookRepository persistence.BookRepository, commandBus commandbus.CommandBus) gorouter.Router {
	router.GET("/book", handlers.BuildListBooksHandler(bookRepository))
	router.GET("/book/{id}", handlers.BuildGetBookHandler(bookRepository))
	router.POST("/dispatch/book/{command}", handlers.BuildBookCommandDispatchHandler(commandBus))
	return router
}
