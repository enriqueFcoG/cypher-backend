package persistence

import (
	"database/sql"
	"enriqueFcoG/cypher/internal/domain"
)

type ModRepoPG struct {
	DB *sql.DB
}

func (r *ModRepoPG) Save(u *domain.Mod) error {
	_, error := r.DB.Exec("INSERT INTO MOD (ID) values($1)", u.ID)
	return error
}

func (r *ModRepoPG) FindByID(id string) (*domain.Mod, error) {
	row := r.DB.QueryRow("SELECT ID, UserName FROM MOD WHERE ID = $1", id)
	u := &domain.Mod{}
	err := row.Scan(&u.ID, &u.ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}
