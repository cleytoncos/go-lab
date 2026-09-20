package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"go-lab/api/internal/domain"
)

type postgresUserRepository struct {
	pool *pgxpool.Pool
}

// NewPostgresUserRepository cria um repositório de usuários backed pelo PostgreSQL.
func NewPostgresUserRepository(pool *pgxpool.Pool) UserRepository {
	return &postgresUserRepository{pool: pool}
}

func (r *postgresUserRepository) FindAll() ([]*domain.User, error) {
	rows, err := r.pool.Query(context.Background(),
		"SELECT id, name, email FROM users ORDER BY id",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		u := &domain.User{}
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *postgresUserRepository) FindByID(id int) (*domain.User, error) {
	u := &domain.User{}
	err := r.pool.QueryRow(context.Background(),
		"SELECT id, name, email FROM users WHERE id = $1", id,
	).Scan(&u.ID, &u.Name, &u.Email)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *postgresUserRepository) Save(user *domain.User) error {
	return r.pool.QueryRow(context.Background(),
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
		user.Name, user.Email,
	).Scan(&user.ID)
}
