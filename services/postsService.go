package services

import (
	"awesomeProject/constants"
	"awesomeProject/models"
	"awesomeProject/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type PostsService struct {
	utils utils.Utils
}

// GetPostsFromEndpoint will fetch all posts from given endpoint
func (service *PostsService) GetPostsFromEndpoint(endpoint string) ([]models.Post, error) {
	dataFromEndpoint, err := http.Get(endpoint)
	if err != nil {
		log.Println("Error: ", err.Error())
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error: ", err.Error())
			return
		}
	}(dataFromEndpoint.Body)
	var posts []models.Post
	if err = json.NewDecoder(dataFromEndpoint.Body).Decode(&posts); err != nil {
		log.Println("Error: ", err.Error())
		return nil, err
	}

	return posts, nil
}

// GetPostById will fetch all posts from given endpoint for the specified user
func (service *PostsService) GetPostById(userId string) ([]models.Post, error) {
	allPosts, err := service.GetPostsFromEndpoint(constants.Endpoint)
	if err != nil {
		return []models.Post{}, err
	}
	convertedUserId, err := service.utils.GetIntegerValueFromString(userId)
	if err != nil {
		return []models.Post{}, err
	}
	log.Println("convertedUserId: ", convertedUserId)
	var userPosts []models.Post
	for _, post := range allPosts {
		if convertedUserId == post.UserId {
			userPosts = append(userPosts, post)
		}
	}
	if len(userPosts) == 0 {
		return []models.Post{}, constants.ErrNoPostsFound
	}
	return userPosts, nil
}
