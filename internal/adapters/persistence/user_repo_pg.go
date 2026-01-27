package persistence

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepoPG struct {
	pool *pgxpool.Pool
}

func NewUserRepoPG(pool *pgxpool.Pool) *UserRepoPG {
	return &UserRepoPG{pool: pool}
}

func (r *UserRepoPG) Save(ctx context.Context, u *domain.User) error {
	query := `INSERT INTO USERS (UserName) values($1)`
	err := r.pool.QueryRow(ctx, query, u.UserName).Scan(&u.UserName)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepoPG) FindByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	query := `SELECT UserName FROM USERS WHERE ID = $1`
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&user.UserName,
	)

	if err != nil {
		return nil, nil
	}

	return &user, nil
}
