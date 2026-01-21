package ports

import "enriqueFcoG/cypher/internal/domain"

type BeatTypeRepository interface {
	Save(b *domain.BeatType) error
	FindByID(id string) (*domain.BeatType, error)
}
