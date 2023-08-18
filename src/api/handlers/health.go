package handlers

import (
	"net/http"

	"github.com/Moji00f/SimpleProject/api/helper"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

type Header struct {
	UserId  string
	Browser string
}

type Person struct {
	FirstName    string `json:"first_name" binding:"required,alpha,min=4,max=10"`
	LastName     string `json:"last_name" binding:"required,alpha,min=6,max=20"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("working", true, 0))
}

func (h *HealthHandler) HeaderBinder1(c *gin.Context) {
	UserId := c.GetHeader("UserId")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "HeaderBinder1",
		"UserId": UserId,
	}, true, 0))
}

func (h *HealthHandler) HeaderBinder2(c *gin.Context) {
	header := Header{}
	c.BindHeader(&header)
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "HeaderBinder2",
		"header": header,
	}, true, 0))
}

func (h *HealthHandler) QueryBinder1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "QueryBinder1",
		"Id":     id,
		"Name":   name,
	}, true, 0))
}

func (h *HealthHandler) QueryBinder2(c *gin.Context) {
	ids := c.QueryArray("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "QueryBInder2",
		"Ids":    ids,
		"name":   name,
	}, true, 0))
}

func (h *HealthHandler) UriBinder(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "UriBinder",
		"id":     id,
		"name":   name,
	}, true, 0))
}

func (h *HealthHandler) BodyBinder(c *gin.Context) {
	p := Person{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidatioinError(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "BodyBinder",
		"person": p,
	}, true, 0))
}

func (h *HealthHandler) FormBinder(c *gin.Context) {
	p := Person{}
	c.ShouldBindJSON(&p)
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "BodyBinder",
		"person": p,
	}, true, 0))
}

func (h *HealthHandler) FileBinder(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"resutl":   "FileBinder",
		"FileName": file.Filename,
	}, true, 0))
}
