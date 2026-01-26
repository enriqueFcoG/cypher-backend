package persistence

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BeatTypeRepoPG struct {
	pool *pgxpool.Pool
}

func NewBeatTypeRepoPG(pool *pgxpool.Pool) *BeatTypeRepoPG {
	return &BeatTypeRepoPG{pool: pool}
}

func (r *BeatTypeRepoPG) Save(ctx context.Context, bt *domain.BeatType) error {
	query := `INSERT INTO BEATTYPE (Name) values($1) RETURNING Name`
	err := r.pool.QueryRow(ctx, query, bt.Name).Scan(&bt.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *BeatTypeRepoPG) FindByID(ctx context.Context, id string) (*domain.BeatType, error) {
	var beat_type domain.BeatType
	query := `SELECT ID, Name FROM BEATTYPE WHERE ID = $1`
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&beat_type.Name,
	)

	if err != nil {
		return nil, nil
	}

	return &beat_type, nil
}
