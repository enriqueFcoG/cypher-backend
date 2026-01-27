package persistence

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type FormatDetailsRepoPG struct {
	pool *pgxpool.Pool
}

func NewFormatDetailsRepoPG(pool *pgxpool.Pool) *FormatDetailsRepoPG {
	return &FormatDetailsRepoPG{pool: pool}
}

func (r *FormatDetailsRepoPG) Save(ctx context.Context, fd *domain.FormatDetails) error {
	query := `INSERT INTO FORMATDETAILS (ID) values($1) RETURNING ID`
	err := r.pool.QueryRow(ctx, query, fd.ID).Scan(&fd.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *FormatDetailsRepoPG) FindByID(ctx context.Context, id string) (*domain.FormatDetails, error) {
	var format_details domain.FormatDetails
	query := `SELECT ID, UserName FROM FORMATDETAILS WHERE ID = $1`
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&format_details.ID,
	)

	if err != nil {
		return nil, nil
	}

	return &format_details, nil
}
