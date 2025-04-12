package api

import (
	"github.com/Moji00f/SimpleProject/api/router"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		router.Health(health)
	}

	r.Run(":8080")
}
