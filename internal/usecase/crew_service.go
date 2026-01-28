package usecase

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"
	"enriqueFcoG/cypher/internal/ports"
)

type CrewService struct {
	Repo ports.CrewRepository
}

func NewCrewService(repo ports.CrewRepository) *CrewService {
	if repo == nil {
		panic("CrewRepository is nill")
	}
	return &CrewService{Repo: repo}
}

func (s *CrewService) CreateCrew(ctx context.Context, name string) error {
	newCrew := domain.Crew{
		Name: name,
	}

	return s.Repo.Save(ctx, &newCrew)
}

func (s *CrewService) GetCrew(ctx context.Context, id string) (*domain.Crew, error) {
	return s.Repo.FindByID(ctx, id)
}
