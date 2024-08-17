package utils

import (
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
