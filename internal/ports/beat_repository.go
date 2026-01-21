package ports

import "enriqueFcoG/cypher/internal/domain"

type BeatRepository interface {
	Save(b *domain.Beat) error
	FindByID(id string) (*domain.Beat, error)
}
