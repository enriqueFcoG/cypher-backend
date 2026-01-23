package persistence

import (
	"database/sql"
	"enriqueFcoG/cypher/internal/domain"
)

type BeatRepoPG struct {
	DB *sql.DB
}

func (r *BeatRepoPG) Save(u *domain.Beat) error {
	_, error := r.DB.Exec("INSERT INTO BEATTYPE (ID, UserName) values($1,$2)", u.ID, u.Title)
	return error
}

func (r *BeatRepoPG) FindByID(id string) (*domain.Beat, error) {
	row := r.DB.QueryRow("SELECT ID, UserName FROM BEATTYPE WHERE ID = $1", id)
	u := &domain.Beat{}
	err := row.Scan(&u.ID, &u.Title)
	if err != nil {
		return nil, err
	}

	return u, nil
}
