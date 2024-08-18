package constants

import "errors"

var (
	ErrNoPostsFound   = errors.New("error no posts found")
	ErrSendingRequest = errors.New("error sending request")
)
