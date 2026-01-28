package usecase

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"
	"enriqueFcoG/cypher/internal/ports"
)

type ModService struct {
	Repo ports.ModRepository
}

func NewModService(repo ports.ModRepository) *ModService {
	if repo == nil {
		panic("ModRepository is nill")
	}
	return &ModService{Repo: repo}
}
func (s *ModService) CreateMod(ctx context.Context, name string) error {
	newMod := domain.Mod{}
	return s.Repo.Save(ctx, &newMod)
}

func (s *ModService) GetMod(ctx context.Context, id string) (*domain.Mod, error) {
	return s.Repo.FindByID(ctx, id)
}
