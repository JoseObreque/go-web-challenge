package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/server/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	data   []domain.Ticket
}

func NewRouter(engine *gin.Engine, data []domain.Ticket) *Router {
	return &Router{
		engine: engine,
		data:   data,
	}
}

func (r *Router) MapRoutes() {
	repo := tickets.NewRepository(r.data)
	service := tickets.NewService(repo)
	handlers := handler.NewService(service)
	r.engine.GET("/ticket/getByCountry/:dest", handlers.GetTicketsByCountry())
	r.engine.GET("/ticket/getAverage/:dest", handlers.AverageDestination())
}
