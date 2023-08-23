package helper

import (
	"net/http"

	"github.com/Moji00f/SimpleProject/pkg/service_error"
)

var StatusCodeMapping = map[string]int{
	//OTP
	service_error.OptExists:   409,
	service_error.OtpUsed:     409,
	service_error.OtpNotValid: 400,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}

	return value
}
