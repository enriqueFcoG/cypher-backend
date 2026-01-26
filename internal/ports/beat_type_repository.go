package ports

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"
)

type BeatTypeRepository interface {
	Save(ctx context.Context, b *domain.BeatType) error
	FindByID(ctx context.Context, id string) (*domain.BeatType, error)
}
