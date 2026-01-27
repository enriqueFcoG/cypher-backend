package ports

import (
	"context"
	"enriqueFcoG/cypher/internal/domain"
)

type UserRepository interface {
	Save(ctx context.Context, u *domain.User) error
	FindByID(ctx context.Context, id string) (*domain.User, error)
	//Update(u *domain.User) (*domain.User, error)
}
