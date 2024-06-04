package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PongController struct {
}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(c *gin.Context) {
	name := c.Param("name")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"uid":     uid,
	})
}
