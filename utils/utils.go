package utils

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
)

type Utils struct {
}

func (utils *Utils) HandleError(err error, responseWriter http.ResponseWriter) {
	if err != nil {
		log.Println("error due to:", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}

func (utils *Utils) GetIntegerValueFromString(str string) (int, error) {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

func (utils *Utils) BuildRequest(url string, method string, headers map[string]string, body []byte) (http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return http.Request{}, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return *req, nil
}
