package routes

import (
	"github.com/Moji00f/SimpleProject/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	health := handlers.NewHealthHandler()

	r.GET("/", health.Health)
	r.GET("/user/:id", health.UserById)

	r.POST("/binder/header1", health.HeaderBinder1)
	r.POST("/binder/header2", health.HeaderBinder2)

	r.POST("/binder/query1", health.QueryBinder1)
	r.POST("/binder/query2", health.QueryBinder2)

	r.POST("/binder/uri/:id/:name", health.UriBinder)

	r.POST("/binder/body", health.BodyBinder)
	r.POST("/binder/form", health.FormBinder)
	r.POST("/binder/file", health.FileBinder)
}
