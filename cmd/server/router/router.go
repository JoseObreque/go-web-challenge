package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/server/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

/*
The Router struct represents the router of the Gin application.

	Fields:
	- engine (*gin.Engine): A gin.Engine instance (for example: gin.Default).
	- data ([]domain.Ticket): A slice of domain.Ticket instances.
*/
type Router struct {
	engine *gin.Engine
	data   []domain.Ticket
}

/*
The NewRouter function is a constructor of the Router struct.

	Parameters:
	 - engine (*gin.Engine): A gin.Engine instance (for example: gin.Default)
	 - data ([]domain.Ticket): A slice of domain.Ticket objects.
*/
func NewRouter(engine *gin.Engine, data []domain.Ticket) Router {
	return Router{
		engine: engine,
		data:   data,
	}
}

/*
The MapRoutes method defines all the routes associated with domain.Ticket objects. It includes:

  - A route for getting the amount of tickets that have a certain destination.
  - A route for getting the average amount of tickets that have a certain destination.
*/
func (r *Router) MapRoutes() {
	repo := tickets.NewRepository(r.data)
	service := tickets.NewService(repo)
	handlers := handler.NewService(service)

	r.engine.GET("/ticket/getByCountry/:dest", handlers.GetTicketsByCountry())
	r.engine.GET("/ticket/getAverage/:dest", handlers.AverageDestination())
}
