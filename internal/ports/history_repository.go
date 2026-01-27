package ports

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"
)

type HistoryRepository interface {
	Save(ctx context.Context, h *domain.History) error
	FindByID(ctx context.Context, id string) (*domain.History, error)
}
