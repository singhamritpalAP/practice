package util

import (
	"log"
	"net/http"
)

func HandleError(err error, responseWriter http.ResponseWriter) {
	if err != nil {
		log.Println("error due to:", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}
