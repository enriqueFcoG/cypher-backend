package ports

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"
)

type ModRepository interface {
	Save(ctx context.Context, m *domain.Mod) error
	FindByID(ctx context.Context, id string) (*domain.Mod, error)
}
