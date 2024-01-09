package handlers

import (
	"go-Expense/pkg/models"
	"go-Expense/pkg/response"
	services "go-Expense/pkg/usecase/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase services.UserUsecase
}

func NewUserHandler(userUsecase services.UserUsecase) UserHandler {
	return UserHandler{
		userUsecase: userUsecase,
	}
}
func (usrH UserHandler) AddSection(c *gin.Context) {
	var section models.Section
	if err := c.BindJSON(section); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "adding section failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
	}
	add, addErr := usrH.userUsecase.AddSection(section.UserName, section.Buget, section.Description, section.CreatedAt)
	if addErr != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "adding failed", nil, addErr.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	success := response.ClientResponse(http.StatusOK, "section added successfully", add, nil)
	c.JSON(http.StatusOK, success)
}
