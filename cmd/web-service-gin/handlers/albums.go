package handlers

import (
	"example.com/web-service-gin/DummyData"
	"example.com/web-service-gin/constants"
	"example.com/web-service-gin/models"
	"example.com/web-service-gin/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service services.PostsService
}

// GetAlbums responds with the list of all albums as JSON.
func (handler *Handler) GetAlbums(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, DummyData.Albums)
}

func (handler *Handler) PostAlbums(ctx *gin.Context) {
	var newAlbum models.Album

	//if err := json.NewDecoder(ctx.Request.Body).Decode(&newAlbum); err != nil {
	//	return
	//}
	if err := ctx.ShouldBindJSON(&newAlbum); err != nil {
		return
	}
	DummyData.Albums = append(DummyData.Albums, newAlbum)
	ctx.JSON(http.StatusOK, newAlbum)
}

func (handler *Handler) GetAllPosts(ctx *gin.Context) {
	post, err := handler.service.GetPostsFromEndpoint(constants.Endpoint)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, post)
}

func (handler *Handler) GetPostById(ctx *gin.Context) {
	id := ctx.Param("id")
	post, err := handler.service.GetPostById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, post)
}
