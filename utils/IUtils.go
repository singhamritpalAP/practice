package utils

import "net/http"

type IUtils interface {
	HandleError(err error, responseWriter http.ResponseWriter)
	GetIntegerValueFromString(str string) (int, error)
	BuildRequest(url string, method string, headers map[string]string, body string) (http.Request, error)
}
