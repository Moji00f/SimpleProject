package handlers

import (
	"net/http"

	"github.com/Moji00f/SimpleProject/api/dto"
	"github.com/Moji00f/SimpleProject/api/helper"
	"github.com/Moji00f/SimpleProject/config"
	"github.com/Moji00f/SimpleProject/services"
	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	service *services.UserService
}

func NewUsersHandler(cfg *config.Config) *UsersHandler {
	service := services.NewUserService(cfg)

	return &UsersHandler{service: service}
}

// SendOtp godoc
// @Summary Send otp to user
// @Description Send otp to user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/send-otp [post]
// @Security AuthBearer
func (h *UsersHandler) SendOtp(c *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidatioinError(nil, false, -1, err))
		return
	}

	err = h.service.SendOtp(req)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err),
		)

		return
	}

	//Call Internal SMS Service

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))
}
