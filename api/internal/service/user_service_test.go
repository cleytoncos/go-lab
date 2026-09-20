package service

import (
	"testing"

	"go-lab/api/internal/domain"
	"go-lab/api/internal/repository"
)

func newService() *UserService {
	return NewUserService(repository.NewInMemoryUserRepository())
}

// --- Create ---

func TestCreate_Success(t *testing.T) {
	svc := newService()

	user, err := svc.Create(&domain.CreateUserRequest{Name: "João", Email: "joao@email.com"})
	if err != nil {
		t.Fatalf("esperava sem erro, obteve: %v", err)
	}
	if user.ID == 0 {
		t.Error("esperava ID preenchido após criação")
	}
	if user.Name != "João" {
		t.Errorf("esperava nome 'João', obteve '%s'", user.Name)
	}
}

func TestCreate_MissingName(t *testing.T) {
	svc := newService()

	_, err := svc.Create(&domain.CreateUserRequest{Name: "", Email: "joao@email.com"})
	if err == nil {
		t.Fatal("esperava erro por nome vazio, obteve nil")
	}
	if err.Error() != "nome é obrigatório" {
		t.Errorf("mensagem de erro inesperada: %s", err.Error())
	}
}

func TestCreate_MissingEmail(t *testing.T) {
	svc := newService()

	_, err := svc.Create(&domain.CreateUserRequest{Name: "João", Email: ""})
	if err == nil {
		t.Fatal("esperava erro por email vazio, obteve nil")
	}
	if err.Error() != "email é obrigatório" {
		t.Errorf("mensagem de erro inesperada: %s", err.Error())
	}
}

// --- GetAll ---

func TestGetAll_Empty(t *testing.T) {
	svc := newService()

	users, err := svc.GetAll()
	if err != nil {
		t.Fatalf("esperava sem erro, obteve: %v", err)
	}
	if len(users) != 0 {
		t.Errorf("esperava lista vazia, obteve %d usuários", len(users))
	}
}

func TestGetAll_WithUsers(t *testing.T) {
	svc := newService()

	svc.Create(&domain.CreateUserRequest{Name: "João", Email: "joao@email.com"})
	svc.Create(&domain.CreateUserRequest{Name: "Maria", Email: "maria@email.com"})

	users, err := svc.GetAll()
	if err != nil {
		t.Fatalf("esperava sem erro, obteve: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("esperava 2 usuários, obteve %d", len(users))
	}
}

// --- GetByID ---

func TestGetByID_Found(t *testing.T) {
	svc := newService()

	created, _ := svc.Create(&domain.CreateUserRequest{Name: "João", Email: "joao@email.com"})

	user, err := svc.GetByID(created.ID)
	if err != nil {
		t.Fatalf("esperava sem erro, obteve: %v", err)
	}
	if user.ID != created.ID {
		t.Errorf("esperava ID %d, obteve %d", created.ID, user.ID)
	}
}

func TestGetByID_NotFound(t *testing.T) {
	svc := newService()

	_, err := svc.GetByID(999)
	if err == nil {
		t.Fatal("esperava erro para ID inexistente, obteve nil")
	}
	if err.Error() != "usuário não encontrado" {
		t.Errorf("mensagem de erro inesperada: %s", err.Error())
	}
}
