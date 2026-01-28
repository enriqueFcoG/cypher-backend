package usecase

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"
	"enriqueFcoG/cypher/internal/ports"
)

type BeatTypeService struct {
	Repo ports.BeatTypeRepository
}

func NewBeatTypeService(repo ports.BeatTypeRepository) *BeatTypeService {
	if repo == nil {
		panic("BeatTypeRepository is nill")
	}
	return &BeatTypeService{Repo: repo}
}

func (s *BeatTypeService) CreateBeatType(ctx context.Context, title string) error {
	beatType := &domain.BeatType{
		Name: title,
	}

	return s.Repo.Save(ctx, beatType)
}

func (s *BeatTypeService) GetBeatType(ctx context.Context, id string) (*domain.BeatType, error) {
	return s.Repo.FindByID(ctx, id)
}
