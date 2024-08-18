package handlers

import (
	"awesomeProject/DummyData"
	"awesomeProject/models"
	"awesomeProject/services"
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
