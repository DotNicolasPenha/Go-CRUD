package post

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/common/logger"
)

type Service struct {
	repository Repository
}

func NewService(r *Repository) *Service {
	if r == nil {
		logger.Fatal("repository of post service is nil")
	}
	return &Service{repository: *r}
}
func (s *Service) AddPost(post CreatePostDTO) error {
	if post.AuthorID.UUID.IsNil() {
		return ErrAuthorIdIsNil
	}
	if post.Body == "" {
		return ErrBodyIsNil
	}
	if err := s.repository.Insert(post); err != nil {
		return err
	}
	return nil
}
func (s *Service) GetPosts() ([]Post, error) {
	posts, err := s.repository.FindMany()
	return posts, err
}
func (s *Service) GetOnePost(id string) (*Post, error) {
	if id == "" {
		return nil, ErrIdParamIsNil
	}
	post, err := s.repository.FindOne(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *Service) RemoveOne(id string) error {
	if id == "" {
		return ErrIdParamIsNil
	}
	return s.repository.DeleteOne(id)
}
