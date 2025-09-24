package httphandlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bharathreddy-97/URLShortner/models"
	"github.com/bharathreddy-97/URLShortner/utils"
)

func ShortenURL(rw http.ResponseWriter, r *http.Request) {
	responseData, err := io.ReadAll(r.Body)
	if err != nil {
		errBody, statusCode := utils.HandleError(utils.E001)
		rw.Write(errBody)
		rw.WriteHeader(statusCode)
		return
	}
	defer r.Body.Close()

	var requestBody models.AddURLRequest
	err = json.Unmarshal(responseData, &requestBody)
	if err != nil {
		errBody, statusCode := utils.HandleError(utils.E002)
		rw.Write(errBody)
		rw.WriteHeader(statusCode)
		return
	}

	// TODO: Logic to shorten the url

	response := &models.GeneralResponse{
		Data: models.AddURLResponse{
			URL: requestBody.URL,
		},
		Status: "OK",
	}
	responseBody, _ := json.Marshal(&response)
	rw.Write(responseBody)
	rw.WriteHeader(200)
}
