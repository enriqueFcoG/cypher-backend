package persistence

import (
	"database/sql"
	"enriqueFcoG/cypher/internal/domain"
)

type FormatDetailsRepoPG struct {
	DB *sql.DB
}

func (r *FormatDetailsRepoPG) Save(u *domain.FormatDetails) error {
	_, error := r.DB.Exec("INSERT INTO FORMATDETAILS (ID) values($1)", u.ID)
	return error
}

func (r *FormatDetailsRepoPG) FindByID(id string) (*domain.FormatDetails, error) {
	row := r.DB.QueryRow("SELECT ID, UserName FROM FORMATDETAILS WHERE ID = $1", id)
	u := &domain.FormatDetails{}
	err := row.Scan(&u.ID, &u.IDBeat)
	if err != nil {
		return nil, err
	}

	return u, nil
}
