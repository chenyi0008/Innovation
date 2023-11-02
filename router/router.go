package router

import (
	"Innovation/middleware"
	"Innovation/service"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func NewRouter() *gin.Engine {
	r = gin.Default()

	r.Group("/user").
		POST("/login", service.LoginService).
		POST("/register", service.RegisterService)

	r.Group("/inspects", middleware.AuthMiddleware()).
		POST("", service.InspectCreate).
		GET("", service.InspectGetAll).
		PUT("", service.InspectUpdate).
		DELETE("/:id", service.InspectDeleteById)

	return r
}
