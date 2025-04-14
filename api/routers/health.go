package routers

import (
	"github.com/Moji00f/SimpleProject/api/handler"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handler.NewHealthHandler()

	r.GET("/", handler.Health)
}
