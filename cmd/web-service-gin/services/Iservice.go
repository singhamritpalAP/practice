package services

import "example.com/web-service-gin/models"

type IService interface {
	GetPostsFromEndpoint(endpoint string) ([]models.Post, error)
	GetPostById(userId string) ([]models.Post, error)
}
