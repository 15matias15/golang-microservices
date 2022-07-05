package utils

type ApplicationErrors struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
	Code       string `json:"code"`
}
