package api

import (
	"fmt"
	"github.com/Moji00f/SimpleProject/api/middleware"
	"github.com/Moji00f/SimpleProject/api/routers"
	"github.com/Moji00f/SimpleProject/docs"

	"github.com/Moji00f/SimpleProject/api/validation"
	"github.com/Moji00f/SimpleProject/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	r.Use(middleware.DefaultStructuredLogger(cfg))
	r.Use(gin.Logger(), gin.Recovery())

	RegisterValidator()
	RegisterRouter(r, cfg)
	RegisterSwagger(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
}

func RegisterValidator() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		if err != nil {
			log.Print(err.Error())
			return
		}
	}
}

func RegisterRouter(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		users := v1.Group("/users")
		test_health := v1.Group("/test")
		routers.Health(health)
		routers.User(users, cfg)
		routers.TestRouter(test_health)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.ExternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
