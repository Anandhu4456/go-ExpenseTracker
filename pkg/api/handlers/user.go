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

// @Summary Add section
// @Description Add a section for a user
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.Section true "User section information"
// @Success 200 {object} response.Response{} "Successfully added section"
// @Failure 400 {object} response.Response{} "Invalid request payload"
// @Router /users/addsection [post]
func (usrH UserHandler) AddSection(c *gin.Context) {
	var section models.Section
	if err := c.BindJSON(&section); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "adding section failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	add, addErr := usrH.userUsecase.AddSection(section.UserName, section.Buget, section.Description)
	if addErr != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "adding failed", nil, addErr.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	c.Header("Content-Type", "application/json")

	success := response.ClientResponse(http.StatusOK, "section added successfully", add, nil)
	c.JSON(http.StatusOK, success)
}

func (usrH UserHandler) DeleteSection(c *gin.Context) {
	hId := c.Param("id")
	if err := usrH.userUsecase.DeleteSection(hId); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "section deleting failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	success := response.ClientResponse(http.StatusOK, "section deleted", nil, nil)
	c.JSON(http.StatusOK, success)
}
