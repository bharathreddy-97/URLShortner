package utils

const (
	E001 = "E001"
	E002 = "E002"
)

var errorDetailsMap = map[string]map[string]interface{}{
	E001: {
		"code":          "E001",
		"errorMessage":  "error reading request body from the request",
		"statusMessage": "Bad Request",
		"statusCode":    400,
	},

	E002: {
		"code":          "E001",
		"errorMessage":  "error unmarshalling request body from the request",
		"statusMessage": "Bad Request",
		"statusCode":    400,
	},
}
