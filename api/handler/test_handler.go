package handler

import (
	"github.com/Moji00f/SimpleProject/api/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestHandler struct {
}

type personData struct {
	Name         string `json:"name" binding:"required,alpha,min=4,max=10"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (t *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("test", true, 0))
}

func (t *TestHandler) BodyBinder(c *gin.Context) {
	p := personData{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, 0, err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "BodyBinder",
		"person": p,
	})
}
