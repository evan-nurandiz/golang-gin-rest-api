package helper

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"erorrs"`
	Data    interface{} `json:"data"`
}

type EmptyObject struct {
}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}

	return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splitedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splitedError,
		Data:    data,
	}

	return res

}
