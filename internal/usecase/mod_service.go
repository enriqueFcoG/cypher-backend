package usecase

import (
	"enriqueFcoG/cypher/internal/domain"
	"enriqueFcoG/cypher/internal/ports"
)

type ModService struct {
	Repo ports.ModRepository
}

func (s *ModService) CreateMod(name string) error {
	newMod := domain.Mod{}
	return s.Repo.Save(&newMod)
}

func (s *ModService) GetMod(id string) (*domain.Mod, error) {
	return s.Repo.FindByID(id)
}
