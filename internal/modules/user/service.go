package post

import "github.com/DotNicolasPenha/Posts-CRUD/internal/common/logger"

type Service struct {
	repository Repository
}

func NewService(repository *Repository) *Service {
	if repository == nil {
		logger.Fatal("repository of user service is nil")
	}
	return &Service{
		repository: *repository,
	}
}
