package ports

import "enriqueFcoG/cypher/internal/domain"

type ModRepository interface {
	Save(m *domain.Mod) error
	FindByID(id string) (*domain.Mod, error)
}
