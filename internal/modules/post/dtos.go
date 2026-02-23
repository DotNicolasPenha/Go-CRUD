package post

import uuid "github.com/jackc/pgtype/ext/gofrs-uuid"

type CreatePostDTO struct {
	AuthorID uuid.UUID `json:"author_id"`
	Username string    `json:"username"`
	Body     string    `json:"body"`
}
