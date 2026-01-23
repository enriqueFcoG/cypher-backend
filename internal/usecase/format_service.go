package usecase

import (
	"enriqueFcoG/cypher/internal/domain"
	"enriqueFcoG/cypher/internal/ports"
)

type FormatService struct {
	Repo ports.FormatRepository
}

func (s *FormatService) CreateFormat(name string) error {
	newFormat := domain.Format{}

	return s.Repo.Save(&newFormat)
}

func (s *FormatService) GetFormat(id string) (*domain.Format, error) {
	return s.Repo.FindByID(id)
}
