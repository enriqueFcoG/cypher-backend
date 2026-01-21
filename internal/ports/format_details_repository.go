package ports

import "enriqueFcoG/cypher/internal/domain"

type FormatDetailsRepository interface {
	Save(f *domain.FormatDetails) error
	FindByID(id string) (*domain.FormatDetails, error)
}
