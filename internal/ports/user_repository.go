package ports

import "enriqueFcoG/cypher/internal/domain"

type UserRepository interface {
	Save(u *domain.User) error
	FindByID(id string) (*domain.User, error)
	//Update(u *domain.User) (*domain.User, error)
}
