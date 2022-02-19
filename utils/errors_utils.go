package utils

type ApplicationError struct {
	Message    string `json:"messsage"`
	StatusCode int    `json:"status"`
	Code       string `json:"code"`
}
