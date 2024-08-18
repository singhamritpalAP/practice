package handlers

import (
	"awesomeProject/constants"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *Handler) GetAllPosts(ctx *gin.Context) {
	post, err := handler.service.GetPostsFromEndpoint(constants.Endpoint)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, post)
}

func (handler *Handler) GetPostById(ctx *gin.Context) {
	id := ctx.Param("id")
	post, err := handler.service.GetPostById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, post)
}

func (handler *Handler) PostContent(ctx *gin.Context) {
	var post models.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	postedContent, err := handler.service.PostContentToEndpoint(post, constants.Endpoint)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, postedContent)
}
