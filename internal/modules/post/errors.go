package post

import "errors"

var (
	ErrNotFound      = errors.New("post not found.")
	ErrAuthorIdIsNil = errors.New("post author id is null.")
	ErrBodyIsNil     = errors.New("post body is null.")
	ErrIdParamIsNil  = errors.New("post id param is undefined.")
)
