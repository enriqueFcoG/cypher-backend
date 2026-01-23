package persistence

import (
	"database/sql"
	"enriqueFcoG/cypher/internal/domain"
)

type BeatTypeRepoPG struct {
	DB *sql.DB
}

func (r *BeatTypeRepoPG) Save(u *domain.BeatType) error {
	_, error := r.DB.Exec("INSERT INTO BEATTYPE (ID, UserName) values($1,$2)", u.ID, u.Name)
	return error
}

func (r *BeatTypeRepoPG) FindByID(id string) (*domain.BeatType, error) {
	row := r.DB.QueryRow("SELECT ID, UserName FROM BEATTYPE WHERE ID = $1", id)
	u := &domain.BeatType{}
	err := row.Scan(&u.ID, &u.Name)
	if err != nil {
		return nil, err
	}

	return u, nil
}
