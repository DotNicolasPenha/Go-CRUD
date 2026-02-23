package user

import "errors"

var (
	ErrNotFound      = errors.New("user not found.")
	ErrNameUserIsNil = errors.New("user name is null.")
	ErrBioIsNil      = errors.New("user bio is null.")
)
