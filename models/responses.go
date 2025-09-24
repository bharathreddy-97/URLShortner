package models

type GeneralResponse struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}

type ErrorResponse struct {
	ErrorCode    string `json:"code"`
	ErrorMessage string `json:"error"`
}

type AddURLResponse struct {
	URL string `json:"url"`
}
