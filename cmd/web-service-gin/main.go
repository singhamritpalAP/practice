package main

import (
	"example.com/web-service-gin/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	handler := handlers.Handler{}
	router.GET("/albums", handler.GetAlbums)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}
