package user

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/common/logger"
)

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
func (s *Service) AddUser(createUserDto CreateUserDTO) error {
	if createUserDto.Username == "" {
		return ErrNameUserIsNil
	}
	if createUserDto.Bio == "" {
		return ErrBioIsNil
	}
	if createUserDto.PasswordHash == "" {
		return ErrPasswordIsNil
	}

	return s.repository.Insert(createUserDto)
}

func (s *Service) GetUsers() ([]User, error) {
	return s.repository.FindMany()
}
