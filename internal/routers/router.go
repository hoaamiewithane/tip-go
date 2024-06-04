package routers

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping/:name", controller.NewPongController().Pong)
		v1.GET("/user/1", controller.NewUserController().GetUserById)
	}
	return r
}
