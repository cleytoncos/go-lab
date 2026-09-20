package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-lab/api/internal/domain"
	"go-lab/api/internal/repository"
	"go-lab/api/internal/service"
)

func newHandler() *UserHandler {
	repo := repository.NewInMemoryUserRepository()
	svc := service.NewUserService(repo)
	return NewUserHandler(svc)
}

// --- GET /users ---

func TestGetAll_Empty(t *testing.T) {
	h := newHandler()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	h.GetAll(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("esperava status 200, obteve %d", w.Code)
	}
}

func TestGetAll_WithUsers(t *testing.T) {
	h := newHandler()

	// cria um usuário diretamente via handler
	body := `{"name":"João","email":"joao@email.com"}`
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	h.Create(w_rec(), req)

	req = httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	h.GetAll(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("esperava status 200, obteve %d", w.Code)
	}

	var users []domain.User
	if err := json.NewDecoder(w.Body).Decode(&users); err != nil {
		t.Fatalf("erro ao decodificar resposta: %v", err)
	}
	if len(users) != 1 {
		t.Errorf("esperava 1 usuário, obteve %d", len(users))
	}
}

// --- POST /users ---

func TestCreate_Success(t *testing.T) {
	h := newHandler()

	body := `{"name":"João","email":"joao@email.com"}`
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Create(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("esperava status 201, obteve %d", w.Code)
	}

	var user domain.User
	if err := json.NewDecoder(w.Body).Decode(&user); err != nil {
		t.Fatalf("erro ao decodificar resposta: %v", err)
	}
	if user.ID == 0 {
		t.Error("esperava ID preenchido na resposta")
	}
	if user.Name != "João" {
		t.Errorf("esperava nome 'João', obteve '%s'", user.Name)
	}
}

func TestCreate_InvalidPayload(t *testing.T) {
	h := newHandler()

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString("not-json"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Create(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("esperava status 400, obteve %d", w.Code)
	}
}

func TestCreate_MissingName(t *testing.T) {
	h := newHandler()

	body := `{"name":"","email":"joao@email.com"}`
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Create(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("esperava status 422, obteve %d", w.Code)
	}
}

// --- GET /users/{id} ---

func TestGetByID_Found(t *testing.T) {
	h := newHandler()

	// cria usuário primeiro
	body := `{"name":"João","email":"joao@email.com"}`
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	wCreate := httptest.NewRecorder()
	h.Create(wCreate, req)

	var created domain.User
	if err := json.NewDecoder(wCreate.Body).Decode(&created); err != nil {
		t.Fatalf("erro ao decodificar resposta: %v", err)
	}

	req = httptest.NewRequest(http.MethodGet, "/users/1", nil)
	w := httptest.NewRecorder()
	h.GetByID(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("esperava status 200, obteve %d", w.Code)
	}

	var user domain.User
	if err := json.NewDecoder(w.Body).Decode(&user); err != nil {
		t.Fatalf("erro ao decodificar resposta: %v", err)
	}
	if user.ID != created.ID {
		t.Errorf("esperava ID %d, obteve %d", created.ID, user.ID)
	}
}

func TestGetByID_InvalidID(t *testing.T) {
	h := newHandler()

	req := httptest.NewRequest(http.MethodGet, "/users/abc", nil)
	w := httptest.NewRecorder()

	h.GetByID(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("esperava status 400, obteve %d", w.Code)
	}
}

func TestGetByID_NotFound(t *testing.T) {
	h := newHandler()

	req := httptest.NewRequest(http.MethodGet, "/users/999", nil)
	w := httptest.NewRecorder()

	h.GetByID(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("esperava status 404, obteve %d", w.Code)
	}
}

// helper para descartar respostas intermediárias
func w_rec() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}
