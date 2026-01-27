package persistence

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type FormatRepoPG struct {
	pool *pgxpool.Pool
}

func NewFormatRepoPG(pool *pgxpool.Pool) *FormatRepoPG {
	return &FormatRepoPG{pool: pool}
}

func (r *FormatRepoPG) Save(ctx context.Context, f *domain.Format) error {
	query := `INSERT INTO FORMAT (Title) values($1)`
	err := r.pool.QueryRow(ctx, query, f.Title).Scan(&f.Title)
	if err != nil {
		return err
	}
	return nil
}

func (r *FormatRepoPG) FindByID(ctx context.Context, id string) (*domain.Format, error) {
	var format domain.Format
	query := `SELECT ID, UserName FROM FORMAT WHERE ID = $1`
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&format.Title,
	)

	if err != nil {
		return nil, nil
	}

	return &format, nil
}
