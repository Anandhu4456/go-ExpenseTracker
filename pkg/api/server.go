package api

import (
	"go-Expense/pkg/api/handlers"
	"go-Expense/pkg/routes"

	"github.com/gin-gonic/gin"
	// _"go-Expense/cmd/api/docs"
	// swaggerfiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	server *gin.Engine
}

func NewServer(userHandler handlers.UserHandler) *Server {
	engine := gin.New()
	engine.Use(gin.Logger())

	// engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.UserRoutes(engine.Group("/users"), userHandler)
	return &Server{
		server: engine,
	}
}

func (ss *Server) StartServer() {
	ss.server.Run(":8000")
}
