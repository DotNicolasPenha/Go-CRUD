package user

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/common/logger"
	"github.com/DotNicolasPenha/Posts-CRUD/pkg/hasher"
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
	password, err := hasher.HashPassword(createUserDto.PasswordHash)
	if err != nil {
		return err
	}
	createUserDto.PasswordHash = password
	return s.repository.Insert(createUserDto)
}

func (s *Service) GetUsers() ([]User, error) {
	return s.repository.FindMany()
}
