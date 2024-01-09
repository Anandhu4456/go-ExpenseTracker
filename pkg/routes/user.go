package routes

import (
	
	"go-Expense/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup, userHandler handlers.UserHandler){

	forUsers:=engine.Group("/users")
	{
		// user routes will go here
		forUsers.POST("/addsection",userHandler.AddSection)
	}
}