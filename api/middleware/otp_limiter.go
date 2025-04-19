package middleware

import (
	"errors"
	"github.com/Moji00f/SimpleProject/api/helper"
	"github.com/Moji00f/SimpleProject/config"
	"github.com/Moji00f/SimpleProject/pkg/limiter"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net"
	"net/http"
	"time"
)

func OtpLimiter(cfg *config.Config) gin.HandlerFunc {
	var ipRateLimiter = limiter.NewIPRateLimiter(rate.Every(cfg.Otp.Limiter*time.Second), 1)
	return func(c *gin.Context) {
		getLimiter := ipRateLimiter.GetLimiter(getIP(c.Request.RemoteAddr))
		if !getLimiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false, helper.OtpLimiterError, errors.New("not allowed")))
			c.Abort()
		}
		c.Next()
	}
}

func getIP(remoteAddr string) string {
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return remoteAddr
	}
	return ip
}
