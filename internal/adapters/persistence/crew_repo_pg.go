package persistence

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CrewRepoPG struct {
	pool *pgxpool.Pool
}

func NewCrewRepoPG(pool *pgxpool.Pool) *CrewRepoPG {
	return &CrewRepoPG{pool: pool}
}

func (r *CrewRepoPG) Save(ctx context.Context, c *domain.Crew) error {
	query := `INSERT INTO CREW (Name) values($1)`
	err := r.pool.QueryRow(ctx, query, c.Name).Scan(&c.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *CrewRepoPG) FindByID(ctx context.Context, id string) (*domain.Crew, error) {
	var crew domain.Crew
	query := `SELECT ID, Name FROM CREW WHERE ID = $1`
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&crew.Name,
	)

	if err != nil {
		return nil, nil
	}

	return &crew, nil
}
