package handlers

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetAlbums(c *gin.Context)
}
