package config

import (
	"os"
	"testing"
)

func setEnv(t *testing.T, vars map[string]string) {
	t.Helper()
	for k, v := range vars {
		os.Setenv(k, v)
	}
	t.Cleanup(func() {
		for k := range vars {
			os.Unsetenv(k)
		}
	})
}

func TestLoad_Defaults(t *testing.T) {
	// garante que não há variáveis definidas
	for _, k := range []string{"PORT", "ENV", "DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"} {
		os.Unsetenv(k)
	}

	cfg := Load()

	if cfg.Port != "8080" {
		t.Errorf("esperava porta '8080', obteve '%s'", cfg.Port)
	}
	if cfg.Env != "development" {
		t.Errorf("esperava env 'development', obteve '%s'", cfg.Env)
	}
}

func TestLoad_CustomPort(t *testing.T) {
	setEnv(t, map[string]string{"PORT": "9090"})

	cfg := Load()

	if cfg.Port != "9090" {
		t.Errorf("esperava porta '9090', obteve '%s'", cfg.Port)
	}
}

func TestBuildDSN(t *testing.T) {
	setEnv(t, map[string]string{
		"DB_USER":     "golab",
		"DB_PASSWORD": "golab123",
		"DB_HOST":     "localhost",
		"DB_PORT":     "5432",
		"DB_NAME":     "golab",
	})

	dsn := buildDSN()
	expected := "postgres://golab:golab123@localhost:5432/golab"

	if dsn != expected {
		t.Errorf("DSN incorreta\nesperava: %s\nobteve:   %s", expected, dsn)
	}
}

func TestBuildDSN_DefaultHostAndPort(t *testing.T) {
	setEnv(t, map[string]string{
		"DB_USER":     "usr",
		"DB_PASSWORD": "pwd",
		"DB_NAME":     "mydb",
	})
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")

	dsn := buildDSN()
	expected := "postgres://usr:pwd@localhost:5432/mydb"

	if dsn != expected {
		t.Errorf("DSN incorreta\nesperava: %s\nobteve:   %s", expected, dsn)
	}
}
