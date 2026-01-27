package usecase

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"
	"enriqueFcoG/cypher/internal/ports"
)

type FormatService struct {
	Repo ports.FormatRepository
}

func (s *FormatService) CreateFormat(ctx context.Context, name string) error {
	newFormat := domain.Format{}

	return s.Repo.Save(ctx, &newFormat)
}

func (s *FormatService) GetFormat(ctx context.Context, id string) (*domain.Format, error) {
	return s.Repo.FindByID(ctx, id)
}
