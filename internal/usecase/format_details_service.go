package usecase

import (
	"enriqueFcoG/cypher/internal/domain"
	"enriqueFcoG/cypher/internal/ports"
)

type FormatDetailsService struct {
	Repo ports.FormatDetailsRepository
}

func (s *FormatDetailsService) CreateFormatDetails(name string) error {
	newFormatDetails := domain.FormatDetails{}
	return s.Repo.Save(&newFormatDetails)
}
func (s *FormatDetailsService) GetFormatDetails(id string) (*domain.FormatDetails, error) {
	return s.Repo.FindByID(id)
}
