package router

import (
	"pacmanzou/HelloWorld/controller"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.POST("/api/auth/register", controller.Register)
}

func Login(r *gin.Engine) {
	r.POST("/api/auth/login", controller.Login)
}
