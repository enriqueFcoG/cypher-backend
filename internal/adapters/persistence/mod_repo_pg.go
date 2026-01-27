package persistence

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ModRepoPG struct {
	pool *pgxpool.Pool
}

func NewModRepoPG(pool *pgxpool.Pool) *ModRepoPG {
	return &ModRepoPG{pool: pool}
}

func (r *ModRepoPG) Save(ctx context.Context, m *domain.Mod) error {
	query := `INSERT INTO MOD (ID) values($1)`
	err := r.pool.QueryRow(ctx, query, m.ID).Scan(&m.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ModRepoPG) FindByID(ctx context.Context, id string) (*domain.Mod, error) {
	var mod domain.Mod
	query := `SELECT ID FROM MOD WHERE ID = $1`
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&mod.ID,
	)

	if err != nil {
		return nil, nil
	}

	return &mod, nil
}
