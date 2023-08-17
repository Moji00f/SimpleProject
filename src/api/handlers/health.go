package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

type Header struct {
	UserId  string
	Browser string
}

type Person struct {
	FirstName string
	LastName  string
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "working",
	})
}

func (h *HealthHandler) HeaderBinder1(c *gin.Context) {
	UserId := c.GetHeader("UserId")
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"UserId": UserId,
	})
}

func (h *HealthHandler) HeaderBinder2(c *gin.Context) {
	header := Header{}
	c.BindHeader(&header)
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder2",
		"header": header,
	})
}

func (h *HealthHandler) QueryBinder1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder1",
		"Id":     id,
		"Name":   name,
	})
}

func (h *HealthHandler) QueryBinder2(c *gin.Context) {
	ids := c.QueryArray("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBInder2",
		"Ids":    ids,
		"name":   name,
	})
}

func (h *HealthHandler) UriBinder(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"result": "UriBinder",
		"id":     id,
		"name":   name,
	})
}

func (h *HealthHandler) BodyBinder(c *gin.Context) {
	p := Person{}
	c.ShouldBindJSON(&p)
	c.JSON(http.StatusOK, gin.H{
		"result": "BodyBinder",
		"person": p,
	})
}

func (h *HealthHandler) FormBinder(c *gin.Context) {
	p := Person{}
	c.ShouldBindJSON(&p)
	c.JSON(http.StatusOK, gin.H{
		"result": "BodyBinder",
		"person": p,
	})
}

func (h *HealthHandler) FileBinder(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"resutl":   "FileBinder",
		"FileName": file.Filename,
	})
}
