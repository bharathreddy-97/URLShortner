package utils

const (
	E001 = "E001"
	E002 = "E002"
	E003 = "E003"
	E004 = "E004"
	E005 = "E005"
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

	E003: {
		"code":          "E003",
		"errorMessage":  "validation error: invalid url",
		"statusMessage": "Bad Request",
		"statusCode":    400,
	},

	E004: {
		"code":          "E004",
		"errorMessage":  "validation error: invalid short url in the request",
		"statusMessage": "Bad Request",
		"statusCode":    400,
	},

	E005: {
		"code":          "E005",
		"errorMessage":  "invalid short url",
		"statusMessage": "Bad Request",
		"statusCode":    400,
	},
}

const allowedLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
