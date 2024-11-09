package handlers

import (
	"awesomeProject/services"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestHandler_GetAllPosts(t *testing.T) {
	type fields struct {
		service services.PostsService
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &Handler{
				service: tt.fields.service,
			}
			handler.GetAllPosts(tt.args.ctx)
		})
	}
}

func TestHandler_GetPostById(t *testing.T) {
	type fields struct {
		service services.PostsService
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &Handler{
				service: tt.fields.service,
			}
			handler.GetPostById(tt.args.ctx)
		})
	}
}

func TestHandler_PostContent(t *testing.T) {
	type fields struct {
		service services.PostsService
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &Handler{
				service: tt.fields.service,
			}
			handler.PostContent(tt.args.ctx)
		})
	}
}

func TestHandler_GetAllPosts1(t *testing.T) {
	type fields struct {
		service services.PostsService
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &Handler{
				service: tt.fields.service,
			}
			handler.GetAllPosts(tt.args.ctx)
		})
	}
}