package ports

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"
)

type BeatRepository interface {
	Save(ctx context.Context, b *domain.Beat) error
	FindByID(ctx context.Context, id string) (*domain.Beat, error)
}
