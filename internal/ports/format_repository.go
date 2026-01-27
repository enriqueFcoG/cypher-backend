package ports

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"
)

type FormatRepository interface {
	Save(ctx context.Context, f *domain.Format) error
	FindByID(ctx context.Context, id string) (*domain.Format, error)
}
