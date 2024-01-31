package response

import "strings"

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObject struct{}

func Success(status int, message string, data interface{}) response {
	return response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

func Error(status int, message string, err string, data interface{}) response {
	spError := strings.Split(err, "\n")
	return response{
		Status:  status,
		Message: message,
		Errors:  spError,
		Data:    data,
	}
}
