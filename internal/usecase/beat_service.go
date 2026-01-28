package usecase

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"
	"enriqueFcoG/cypher/internal/ports"
)

type BeatService struct {
	Repo ports.BeatRepository
}

func NewBeatService(repo ports.BeatRepository) *BeatService {
	if repo == nil {
		panic("BeatRepository is nill")
	}
	return &BeatService{Repo: repo}
}

func (s *BeatService) CreateBeat(ctx context.Context, title string) error {
	beat := &domain.Beat{
		Title:     title,
		Producer:  "string",
		Stream:    "string",
		Rithm:     "string",
		CreatedAt: "string",
		IDType:    "string",
	}

	return s.Repo.Save(ctx, beat)
}

func (s *BeatService) GetBeat(ctx context.Context, id string) (*domain.Beat, error) {
	return s.Repo.FindByID(ctx, id)
}
