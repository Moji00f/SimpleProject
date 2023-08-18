package helper

import (
	"github.com/Moji00f/SimpleProject/api/validations"
)

type BaseHttpResponse struct {
	Resule           any                            `json:"result"`
	Success          bool                           `json:"success"`
	ResultCode       int                            `json:"resultCode"`
	ValidationErrors *[]validations.ValidationError `json:"validationErrors"`
	Error            any                            `json:"error"`
}

func GenerateBaseResponse(result any, success bool, resultCode int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Resule:     result,
		Success:    success,
		ResultCode: resultCode,
	}
}

func GenerateBaseResponseWithError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Resule:     result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err,
	}
}

func GenerateBaseResponseWithValidatioinError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Resule:           result,
		Success:          success,
		ResultCode:       resultCode,
		ValidationErrors: validations.GetValidationErrors(err),
	}
}
