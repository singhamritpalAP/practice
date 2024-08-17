package main

import (
	"awesomeProject/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	handler := handlers.Handler{}
	router.GET("/albums", handler.GetAlbums)
	router.POST("/albums", handler.PostAlbums)
	router.GET("/posts", handler.GetAllPosts)
	router.GET("/posts/:id", handler.GetPostById)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}
