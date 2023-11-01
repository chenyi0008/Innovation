package router

import (
	"Innovation/service"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func NewRouter() *gin.Engine {
	r = gin.Default()

	r.Group("/user").
		POST("/login", service.LoginService).
		POST("/register", service.RegisterService)

	return r
}
