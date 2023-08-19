package api

import (
	"fmt"

	"github.com/Moji00f/SimpleProject/api/middlewares"
	"github.com/Moji00f/SimpleProject/api/routes"
	"github.com/Moji00f/SimpleProject/api/validations"
	"github.com/Moji00f/SimpleProject/config"
	"github.com/Moji00f/SimpleProject/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	RegisterValidators()

	r.Use(gin.Logger(), gin.Recovery() /*middlewares.TestMiddleware()*/, middlewares.LimitByRequest())

	RegisterRoutes(r)
	RegisterSwagger(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
	}
}

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		routes.Health(v1)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "Golang Web API"
	docs.SwaggerInfo.Description = "Golang Web API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
