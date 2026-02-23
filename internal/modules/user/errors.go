package user

import "errors"

var (
	ErrNotFound      = errors.New("user not found.")
	ErrAuthorIdIsNil = errors.New("user name is null.")
	ErrBodyIsNil     = errors.New("user bio is null.")
)
