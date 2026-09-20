// Package main é o entrypoint da aplicação.
package main

import (
	"fmt"
	"log"
	"net/http"

	_ "go-lab/api/docs" // gerado pelo swag init
	"go-lab/api/internal/config"
	"go-lab/api/internal/db"
	"go-lab/api/routes"
)

// @title           go-lab API
// @version         1.0
// @description     API de exemplo para estudo de Go bem arquitetado.

// @contact.name    Seu Nome
// @contact.email   seu@email.com

// @license.name    MIT

// @host            localhost:8080
// @BasePath        /
func main() {
	cfg := config.Load()

	pool, err := db.Connect(cfg.DSN)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}
	defer pool.Close()
	log.Println("Conectado ao PostgreSQL")

	router := routes.Register(pool)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Servidor rodando em http://localhost%s", addr)
	log.Printf("Documentação em  http://localhost%s/docs/index.html", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
