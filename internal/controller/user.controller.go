package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce/internal/service"
	"go-ecommerce/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	response.SuccessResponse(c, 20001, []string{"user1", "user2"})
}
