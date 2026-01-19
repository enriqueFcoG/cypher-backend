package persistence

import (
	"database/sql"
	"enriqueFcoG/cypher/internal/domain"
)

type UserRepoPG struct {
	DB *sql.DB
}

func (r *UserRepoPG) Save(u *domain.User) error {
	_, error := r.DB.Exec("INSERT INTO USERS (ID, UserName) values($1,$2)", u.ID, u.UserName)
	return error
}

func (r *UserRepoPG) FindByID(id string) (*domain.User, error) {
	row := r.DB.QueryRow("SELECT ID, UserName FROM USERS WHERE ID = $1", id)
	u := &domain.User{}
	err := row.Scan(&u.ID, &u.UserName)
	if err != nil {
		return nil, err
	}

	return u, nil
}
