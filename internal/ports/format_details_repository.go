package ports

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"
)

type FormatDetailsRepository interface {
	Save(ctx context.Context, f *domain.FormatDetails) error
	FindByID(ctx context.Context, id string) (*domain.FormatDetails, error)
}
