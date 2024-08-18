package handlers

import "github.com/gin-gonic/gin"

type IHandler interface {
	GetAlbums(ctx *gin.Context)
	GetAllPosts(ctx *gin.Context)
	GetPostById(ctx *gin.Context)
	PostContent(ctx *gin.Context)
}
