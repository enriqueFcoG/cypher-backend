package persistence

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HistoryRepoPG struct {
	pool *pgxpool.Pool
}

func NewHistoryRepoPG(pool *pgxpool.Pool) *HistoryRepoPG {
	return &HistoryRepoPG{pool: pool}
}

func (r *HistoryRepoPG) Save(ctx context.Context, h *domain.History) error {
	query := `INSERT INTO HISTORY (ID) values($1)`
	err := r.pool.QueryRow(ctx, query, h.ID).Scan(&h.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *HistoryRepoPG) FindByID(ctx context.Context, id string) (*domain.History, error) {
	var history domain.History
	query := `SELECT ID, UserName FROM HISTORY WHERE ID = $1`
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&history.ID,
	)

	if err != nil {
		return nil, nil
	}

	return &history, nil
}
