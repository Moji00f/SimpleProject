package routers

import (
	"github.com/Moji00f/SimpleProject/api/handler"
	"github.com/Moji00f/SimpleProject/api/middleware"
	"github.com/Moji00f/SimpleProject/config"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewUserHandler(cfg)

	router.POST("/send-otp", middleware.OtpLimiter(cfg), h.SendOtp)
}
