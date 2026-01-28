package usecase

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"
	"enriqueFcoG/cypher/internal/ports"
)

type HistoryService struct {
	Repo ports.HistoryRepository
}

func NewHistoryService(repo ports.HistoryRepository) *HistoryService {
	if repo == nil {
		panic("HistoryRepository is nill")
	}
	return &HistoryService{Repo: repo}
}

func (s *HistoryService) CreateHistory(ctx context.Context, name string) error {
	newHistory := domain.History{}

	return s.Repo.Save(ctx, &newHistory)
}

func (s *HistoryService) GetHistory(ctx context.Context, id string) (*domain.History, error) {
	return s.Repo.FindByID(ctx, id)
}
