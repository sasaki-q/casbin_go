package pkg

import "errors"

var (
	ErrContextValueNotFound  = errors.New("context value not found")
	ErrHeaderValueNotFound   = errors.New("uid not found")
	ErrPathParameterNotFound = errors.New("path parameter not found")
	ErrNotAllowed            = errors.New("not allowed")
	ErrUserNotFound          = errors.New("user not found")
)
