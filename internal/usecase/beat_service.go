package usecase

import (
	"enriqueFcoG/cypher/internal/domain"
	"enriqueFcoG/cypher/internal/ports"
)

type BeatService struct {
	Repo ports.BeatRepository
}

func (s *BeatService) CreateBeat(title string) error {
	beat := &domain.Beat{
		Title:     title,
		Producer:  "string",
		Stream:    "string",
		Rithm:     "string",
		CreatedAt: "string",
		IDType:    "string",
	}

	return s.Repo.Save(beat)
}

func (s *BeatService) GetBear(id string) (*domain.Beat, error) {
	return s.Repo.FindByID(id)
}
