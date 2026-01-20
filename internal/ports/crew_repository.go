package ports

import "enriqueFcoG/cypher/internal/domain"

type CrewRepository interface {
	Save(c *domain.Crew) error
	FindByID(id string) (*domain.Crew, error)
}
