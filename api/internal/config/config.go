package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config contém todas as configurações da aplicação.
type Config struct {
	Port string
	Env  string
	DSN  string // Data Source Name montada a partir das variáveis de banco
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
		DSN:  buildDSN(),
	}
}

// buildDSN monta a connection string a partir das variáveis individuais de banco.
func buildDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		getEnv("DB_USER", ""),
		getEnv("DB_PASSWORD", ""),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_NAME", ""),
	)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
