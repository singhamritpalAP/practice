package services

import (
	"awesomeProject/models"
)

type IService interface {
	GetPostsFromEndpoint(endpoint string) ([]models.Post, error)
	GetPostById(userId string) ([]models.Post, error)
	PostContentToEndpoint(content models.Post, endPoint string) (models.Post, error)
}
