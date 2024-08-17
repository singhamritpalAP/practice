package handlers

import (
	"example.com/web-service-gin/DummyData"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

// GetAlbums responds with the list of all albums as JSON.
func (handler *Handler) GetAlbums(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, DummyData.Albums)
}
