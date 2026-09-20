package domain

// User representa a entidade de usuário no domínio da aplicação.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUserRequest é o payload esperado para criação de um usuário.
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
