package persistence

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BeatRepoPG struct {
	pool *pgxpool.Pool
}

func NewBeatRepoPG(pool *pgxpool.Pool) *BeatRepoPG {
	return &BeatRepoPG{pool: pool}
}

func (r *BeatRepoPG) Save(ctx context.Context, b *domain.Beat) error {
	query := `INSERT INTO BEAT (Name) values($1) RETURNING Name`
	err := r.pool.QueryRow(ctx, query, b.Type.Name).Scan(&b.Type.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *BeatRepoPG) FindByID(ctx context.Context, id string) (*domain.Beat, error) {
	var beat domain.Beat
	query := `SELECT ID FROM BEAT WHERE ID = $1`
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&beat.ID,
	)

	if err != nil {
		return nil, nil
	}

	return &beat, nil
}
