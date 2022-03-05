package utils

import "strings"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status int, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func BuildErrorResponse(status int, message string, err string) Response {
	splittedError := strings.Split(err, "\n")
	return Response{
		Status:  status,
		Message: message,
		Error:   splittedError,
	}
}
