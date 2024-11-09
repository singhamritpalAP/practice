package services

import (
	"awesomeProject/models"
	"awesomeProject/utils"
	"reflect"
	"testing"
)

func TestPostsService_GetPostsFromEndpoint(t *testing.T) {
	type fields struct {
		utils utils.Utils
	}
	type args struct {
		endpoint string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &PostsService{
				utils: tt.fields.utils,
			}
			got, err := service.GetPostsFromEndpoint(tt.args.endpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPostsFromEndpoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPostsFromEndpoint() got = %v, want %v", got, tt.want)
			}
		})
	}
}
