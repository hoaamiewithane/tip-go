package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	name := c.DefaultQuery("name", "hoaa")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"uid":     uid,
		"user":    []string{"cr7", "messi", "neymar"},
	})
}
