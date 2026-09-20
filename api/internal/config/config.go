package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config contém todas as configurações da aplicação.
type Config struct {
	Port string
	Env  string
	DSN  string // Data Source Name para conexão com banco
}

// Load lê o arquivo .env (se existir) e as variáveis de ambiente,
// retornando uma Config populada com valores padrão como fallback.
func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	return &Config{
		Port: getEnv("PORT", "8080"),
		Env:  getEnv("ENV", "development"),
		DSN:  getEnv("DATABASE_URL", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
