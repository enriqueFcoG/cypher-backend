package usecase

import (
	"enriqueFcoG/cypher/internal/domain"
	"enriqueFcoG/cypher/internal/ports"
)

type HistoryService struct {
	Repo ports.HistoryRepository
}

func (s *HistoryService) CreateHistory(name string) error {
	newHistory := domain.History{}

	return s.Repo.Save(&newHistory)
}

func (s *HistoryService) GetHistory(id string) (*domain.History, error) {
	return s.Repo.FindByID(id)
}
