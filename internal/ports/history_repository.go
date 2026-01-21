package ports

import "enriqueFcoG/cypher/internal/domain"

type HistoryRepository interface {
	Save(h *domain.History) error
	FindByID(id string) (*domain.History, error)
}
