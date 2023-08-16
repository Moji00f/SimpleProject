package routes

import (
	"github.com/Moji00f/SimpleProject/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	health := handlers.NewHealthHandler()

	r.GET("/", health.Health)
}
