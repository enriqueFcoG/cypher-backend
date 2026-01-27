package usecase

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"
	"enriqueFcoG/cypher/internal/ports"
)

type UserService struct {
	Repo ports.UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, id, username string) error {
	user := &domain.User{
		ID:       id,
		UserName: username,
		Photo:    username,
		Respect:  0,
		Email:    username,
		Password: username,
		Online:   true,
		Country:  username,
		State:    username,
		City:     username,
		IDCrew:   username,
	}
	return s.Repo.Save(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	return s.Repo.FindByID(ctx, id)
}
