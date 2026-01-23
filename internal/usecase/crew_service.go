package usecase

import (
	"enriqueFcoG/cypher/internal/domain"
	"enriqueFcoG/cypher/internal/ports"
)

type CrewService struct {
	Repo ports.CrewRepository
}

func (s *CrewService) CreateCrew(name string) error {
	newCrew := domain.Crew{
		Name: name,
	}

	return s.Repo.Save(&newCrew)
}

func (s *CrewService) GetCrew(id string) (*domain.Crew, error) {
	return s.Repo.FindByID(id)
}
