package usecase

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"
	"enriqueFcoG/cypher/internal/ports"
)

type FormatDetailsService struct {
	Repo ports.FormatDetailsRepository
}

func NewFormatDetailsService(repo ports.FormatDetailsRepository) *FormatDetailsService {
	if repo == nil {
		panic("FormatDetailsRepository is nill")
	}
	return &FormatDetailsService{Repo: repo}
}

func (s *FormatDetailsService) CreateFormatDetails(ctx context.Context, name string) error {
	newFormatDetails := domain.FormatDetails{}
	return s.Repo.Save(ctx, &newFormatDetails)
}
func (s *FormatDetailsService) GetFormatDetails(ctx context.Context, id string) (*domain.FormatDetails, error) {
	return s.Repo.FindByID(ctx, id)
}
