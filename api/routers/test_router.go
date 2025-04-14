package routers

import (
	"github.com/Moji00f/SimpleProject/api/handler"
	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	h := handler.NewTestHandler()
	r.GET("/", h.Test)
	r.POST("/binder/body", h.BodyBinder)
}
