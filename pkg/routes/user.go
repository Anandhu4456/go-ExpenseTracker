package routes

import (
	"go-Expense/pkg/api/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup, userHandler handlers.UserHandler) {
	log.Println("setting up user routes...")
	engine.POST("/addsection",userHandler.AddSection)
}
