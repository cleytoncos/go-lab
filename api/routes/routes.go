package routes

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger"

	"go-lab/api/internal/handler"
	"go-lab/api/internal/repository"
	"go-lab/api/internal/service"
)

// Register monta todas as rotas da API e retorna o http.ServeMux pronto para uso.
func Register(pool *pgxpool.Pool) *http.ServeMux {
	// Composição das dependências (Dependency Injection manual)
	userRepo := repository.NewPostgresUserRepository(pool)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	mux := http.NewServeMux()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetAll(w, r)
		case http.MethodPost:
			userHandler.Create(w, r)
		default:
			http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetByID(w, r)
		default:
			http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		}
	})

	// Documentação: http://localhost:8080/docs/index.html
	mux.Handle("/docs/", httpSwagger.Handler(
		httpSwagger.URL("/docs/doc.json"),
	))

	return mux
}
