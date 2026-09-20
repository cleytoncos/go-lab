package repository

import "go-lab/api/internal/domain"

// UserRepository define o contrato de acesso a dados para usuários.
// Qualquer implementação (in-memory, PostgreSQL, etc.) deve satisfazer esta interface.
type UserRepository interface {
	FindAll() ([]*domain.User, error)
	FindByID(id int) (*domain.User, error)
	Save(user *domain.User) error
}

// inMemoryUserRepository é uma implementação simples em memória, útil para estudos e testes.
type inMemoryUserRepository struct {
	store  map[int]*domain.User
	nextID int
}

// NewInMemoryUserRepository cria e retorna um repositório em memória.
func NewInMemoryUserRepository() UserRepository {
	return &inMemoryUserRepository{
		store:  make(map[int]*domain.User),
		nextID: 1,
	}
}

func (r *inMemoryUserRepository) FindAll() ([]*domain.User, error) {
	users := make([]*domain.User, 0, len(r.store))
	for _, u := range r.store {
		users = append(users, u)
	}
	return users, nil
}

func (r *inMemoryUserRepository) FindByID(id int) (*domain.User, error) {
	u, ok := r.store[id]
	if !ok {
		return nil, nil
	}
	return u, nil
}

func (r *inMemoryUserRepository) Save(user *domain.User) error {
	user.ID = r.nextID
	r.store[r.nextID] = user
	r.nextID++
	return nil
}
