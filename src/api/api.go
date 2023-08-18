package api

import (
	"fmt"

	"github.com/Moji00f/SimpleProject/api/middlewares"
	"github.com/Moji00f/SimpleProject/api/routes"
	"github.com/Moji00f/SimpleProject/api/validations"
	"github.com/Moji00f/SimpleProject/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
	}

	r.Use(gin.Logger(), gin.Recovery() /*middlewares.TestMiddleware()*/, middlewares.LimitByRequest())

	v1 := r.Group("/api/v1")
	{
		routes.Health(v1)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
