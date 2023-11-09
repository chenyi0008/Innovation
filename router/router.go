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
		DELETE("/:id", service.InspectDeleteById).
		GET("/:id", service.InspectGetInfoById).
		POST("/bind", service.InspectBindAlarm).
		DELETE("bind", service.InspectUnBindAlarm)

	r.Group("/alarms", middleware.AuthMiddleware()).
		POST("", service.AlarmCreate).
		GET("", service.AlarmGetAll).
		PUT("", service.AlarmUpdate).
		DELETE("/:id", service.AlarmDeleteById)

	r.Group("/histories").
		POST("/alarm", service.AlarmAndUpload).
		GET("/download", service.HistoryDownload)

	return r
}
