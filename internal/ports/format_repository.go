package ports

import "enriqueFcoG/cypher/internal/domain"

type FormatRepository interface {
	Save(f *domain.Format) error
	FindByID(id string) (*domain.Format, error)
}
