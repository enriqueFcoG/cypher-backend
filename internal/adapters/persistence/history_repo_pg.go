package persistence

import (
	"database/sql"
	"enriqueFcoG/cypher/internal/domain"
)

type HistoryRepoPG struct {
	DB *sql.DB
}

func (r *HistoryRepoPG) Save(u *domain.History) error {
	_, error := r.DB.Exec("INSERT INTO HISTORY (ID) values($1)", u.ID)
	return error
}

func (r *HistoryRepoPG) FindByID(id string) (*domain.History, error) {
	row := r.DB.QueryRow("SELECT ID, UserName FROM HISTORY WHERE ID = $1", id)
	u := &domain.History{}
	err := row.Scan(&u.ID, &u.ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}
