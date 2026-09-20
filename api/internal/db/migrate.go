package db

import (
	"embed"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // driver de destino postgres://
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// migrationsFS embute os arquivos .sql da pasta migrations no binário,
// eliminando a dependência de a pasta existir em disco em produção.
//
//go:embed migrations/*.sql
var migrationsFS embed.FS

// RunMigrations aplica todas as migrations pendentes usando o golang-migrate.
// É idempotente: migrations já aplicadas são ignoradas.
func RunMigrations(dsn string) error {
	source, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return fmt.Errorf("erro ao carregar migrations embutidas: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", source, ensureSSLMode(dsn))
	if err != nil {
		return fmt.Errorf("erro ao inicializar migrate: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("erro ao aplicar migrations: %w", err)
	}

	return nil
}

// ensureSSLMode garante que a DSN tenha um sslmode definido. O driver postgres
// do golang-migrate exige SSL por padrão; em ambiente local sem SSL usamos
// sslmode=disable. Se a DSN já define sslmode, ela é preservada.
func ensureSSLMode(dsn string) string {
	if strings.Contains(dsn, "sslmode=") {
		return dsn
	}
	sep := "?"
	if strings.Contains(dsn, "?") {
		sep = "&"
	}
	return dsn + sep + "sslmode=disable"
}
