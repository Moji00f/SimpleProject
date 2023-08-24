package routes

import (
	"github.com/Moji00f/SimpleProject/api/handlers"
	"github.com/Moji00f/SimpleProject/api/middlewares"
	"github.com/Moji00f/SimpleProject/config"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)

	router.POST("/send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
}
