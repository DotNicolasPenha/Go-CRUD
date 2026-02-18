package post

import "errors"

type Service struct {
	repository Repository
}

func NewService(r *Repository) *Service {
	if r == nil {
		panic("repository of post service is nil")
	}
	return &Service{repository: *r}
}
func (s *Service) AddPost(post Post) error {
	if post.Username == "" {
		return errors.New("the username of post is empty")
	}
	if post.Body == "" {
		return errors.New("the body of post is empty")
	}
	if err := s.repository.Insert(post); err != nil {
		return err
	}
	return nil
}
