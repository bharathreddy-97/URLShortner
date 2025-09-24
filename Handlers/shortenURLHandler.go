package httphandlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	persistence "github.com/bharathreddy-97/URLShortner/Persistence"
	"github.com/bharathreddy-97/URLShortner/models"
	"github.com/bharathreddy-97/URLShortner/utils"
)

type Handlers struct {
	MapData *persistence.SMap
}

func (h *Handlers) ShortenURL(rw http.ResponseWriter, r *http.Request) {
	requestBody, errString := parseRequest(r)
	if errString != "" || requestBody == nil {
		handleError(rw, errString)
		return
	}

	_, err := url.ParseRequestURI(requestBody.URL)
	if err != nil {
		handleError(rw, utils.E003)
		return
	}

	shortURL := utils.CreateShortURL(requestBody.URL, h.MapData)

	response := &models.GeneralResponse{
		Data: models.AddURLResponse{
			URL: shortURL,
		},
		Status: "OK",
	}
	responseBody, _ := json.Marshal(&response)
	rw.Write(responseBody)
	rw.WriteHeader(200)
}

func parseRequest(r *http.Request) (*models.AddURLRequest, string) {
	responseData, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, utils.E001
	}
	defer r.Body.Close()

	var requestBody models.AddURLRequest
	err = json.Unmarshal(responseData, &requestBody)
	if err != nil {
		return nil, utils.E002
	}

	return &requestBody, ""
}

func handleError(rw http.ResponseWriter, code string) {
	errBody, statusCode := utils.HandleError(code)
	rw.Write(errBody)
	rw.WriteHeader(statusCode)
}
