package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go-lab/api/internal/domain"
	"go-lab/api/internal/service"
	"go-lab/api/pkg/response"
)

// UserHandler agrupa os handlers HTTP relacionados a usuários.
type UserHandler struct {
	service *service.UserService
}

// NewUserHandler cria um UserHandler com o serviço injetado.
func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{service: svc}
}

// GetAll retorna todos os usuários.
//
// @Summary      Lista todos os usuários
// @Description  Retorna a lista completa de usuários cadastrados
// @Tags         users
// @Produce      json
// @Success      200  {array}   domain.User
// @Failure      500  {object}  map[string]string
// @Router       /users [get]
func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAll()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, users)
}

// GetByID retorna um usuário pelo ID.
//
// @Summary      Busca usuário por ID
// @Description  Retorna um único usuário a partir do ID informado na URL
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "ID do usuário"
// @Success      200  {object}  domain.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [get]
func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "ID inválido")
		return
	}

	user, err := h.service.GetByID(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, user)
}

// Create cria um novo usuário.
//
// @Summary      Cria um novo usuário
// @Description  Recebe nome e email e persiste um novo usuário
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      domain.CreateUserRequest  true  "Dados do usuário"
// @Success      201   {object}  domain.User
// @Failure      400   {object}  map[string]string
// @Failure      422   {object}  map[string]string
// @Router       /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "payload inválido")
		return
	}

	user, err := h.service.Create(&req)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, user)
}
