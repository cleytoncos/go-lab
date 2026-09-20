package service

import (
	"errors"

	"go-lab/api/internal/domain"
	"go-lab/api/internal/repository"
)

// UserService contém as regras de negócio relacionadas a usuários.
type UserService struct {
	repo repository.UserRepository
}

// NewUserService cria um UserService com o repositório injetado.
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]*domain.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetByID(id int) (*domain.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("usuário não encontrado")
	}
	return user, nil
}

func (s *UserService) Create(req *domain.CreateUserRequest) (*domain.User, error) {
	if req.Name == "" {
		return nil, errors.New("nome é obrigatório")
	}
	if req.Email == "" {
		return nil, errors.New("email é obrigatório")
	}

	user := &domain.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := s.repo.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}
