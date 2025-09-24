package utils

import (
	"encoding/json"

	"github.com/bharathreddy-97/URLShortner/models"
)

type errorDetails struct {
	code          string
	message       string
	statusMessage string
	statusCode    int
}

func HandleError(code string) ([]byte, int) {
	errorDetails := getErrorDetails(code)
	errorResponse := &models.ErrorResponse{
		ErrorCode:    errorDetails.code,
		ErrorMessage: errorDetails.message,
	}

	data, _ := json.Marshal(&models.GeneralResponse{
		Data:   errorResponse,
		Status: errorDetails.statusMessage,
	})
	return data, errorDetails.statusCode
}

func getErrorDetails(code string) *errorDetails {
	details := errorDetailsMap[code]
	return &errorDetails{
		code:          details["code"].(string),
		message:       details["errorMessage"].(string),
		statusMessage: details["statusMessage"].(string),
		statusCode:    details["statusCode"].(int),
	}
}
