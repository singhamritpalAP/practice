package main

import (
	"awesomeProject/models"
	"awesomeProject/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

const (
	EndpointUrl = "https://jsonplaceholder.typicode.com/posts"
)

func main() {
	http.HandleFunc("/fetch-post/", handleFunc)
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleFunc(responseWriter http.ResponseWriter, request *http.Request) {
	userIdString := request.URL.Query().Get("userId")
	log.Println("User ID: " + userIdString)
	responseFromEndpoint, err := http.Get(EndpointUrl)
	if err != nil {
		util.HandleError(err, responseWriter)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("error closing response body")
			return
		}
	}(responseFromEndpoint.Body)

	var posts []models.Post
	err = json.NewDecoder(responseFromEndpoint.Body).Decode(&posts)
	if err != nil {
		util.HandleError(err, responseWriter)
		return
	}

	var userPosts []models.Post
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		util.HandleError(fmt.Errorf("invalid userId received"), responseWriter)
		return
	}

	for _, post := range posts {
		if post.UserId == userId {
			userPosts = append(userPosts, post)
		}
	}

	if len(userPosts) == 0 {
		util.HandleError(fmt.Errorf("no posts for requested user"), responseWriter)
		return
	}
	responseWriter.Header().Add("Content-Type", "application/json")
	responseBytes, err := json.Marshal(userPosts)
	if err != nil {
		util.HandleError(err, responseWriter)
		return
	}
	_, err = responseWriter.Write(responseBytes)
	if err != nil {
		util.HandleError(err, responseWriter)
		return
	}

}
