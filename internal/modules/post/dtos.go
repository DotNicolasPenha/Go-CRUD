package post

import uuid "github.com/jackc/pgtype/ext/gofrs-uuid"

type CreatePostDTO struct {
	AuthorID uuid.UUID `json:"author_id"`
	Body     string    `json:"body"`
}
