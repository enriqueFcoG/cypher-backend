package usecase

import (
	"enriqueFcoG/cypher/internal/domain"
	"enriqueFcoG/cypher/internal/ports"
)

type BeatTypeService struct {
	Repo ports.BeatTypeRepository
}

func (s *BeatTypeService) CreateBeatType(title string) error {
	beatType := &domain.BeatType{
		Name: title,
	}

	return s.Repo.Save(beatType)
}

func (s *BeatTypeService) GetBeatType(id string) (*domain.BeatType, error) {
	return s.Repo.FindByID(id)
}
