package api

import (
	"fmt"

	"github.com/Moji00f/SimpleProject/api/routes"
	"github.com/Moji00f/SimpleProject/config"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		routes.Health(v1)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
