package main

import (
	"awesomeProject/models"
	"awesomeProject/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/fetch-post/", handlerFunc)
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerFunc(responseWriter http.ResponseWriter, request *http.Request) {
	userIdString := request.URL.Query().Get("userId")
	log.Println("User ID: " + userIdString)
	responseFromEndpoint, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		utils.HandleError(err, responseWriter)
		return
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println("error closing response body")
			return
		}
	}(responseFromEndpoint.Body)

	var posts []models.Post
	err = json.NewDecoder(responseFromEndpoint.Body).Decode(&posts)
	if err != nil {
		utils.HandleError(err, responseWriter)
		return
	}
	var userPosts []models.Post
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		utils.HandleError(err, responseWriter)
		return
	}

	for _, post := range posts {
		if post.UserId == userId {
			userPosts = append(userPosts, post)
		}
	}

	if len(userPosts) == 0 {
		utils.HandleError(fmt.Errorf(fmt.Sprintf("no posts found for the user with id: %d ", userId)),
			responseWriter)
		return
	}

	responseBody, err := json.Marshal(&userPosts)
	if err != nil {
		utils.HandleError(err, responseWriter)
		return
	}

	responseWriter.Header().Add("Content-Type", "application/json")
	_, err = responseWriter.Write(responseBody)
	if err != nil {
		utils.HandleError(err, responseWriter)
		return
	}
}
