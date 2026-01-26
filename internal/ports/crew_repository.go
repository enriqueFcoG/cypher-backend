package ports

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"
)

type CrewRepository interface {
	Save(ctx context.Context, c *domain.Crew) error
	FindByID(ctx context.Context, id string) (*domain.Crew, error)
}
