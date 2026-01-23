package persistence

import (
	"database/sql"
	"enriqueFcoG/cypher/internal/domain"
)

type FormatRepoPG struct {
	DB *sql.DB
}

func (r *FormatRepoPG) Save(u *domain.Format) error {
	_, error := r.DB.Exec("INSERT INTO FORMATDETAILS (ID) values($1)", u.ID)
	return error
}

func (r *FormatRepoPG) FindByID(id string) (*domain.Format, error) {
	row := r.DB.QueryRow("SELECT ID, UserName FROM FORMATDETAILS WHERE ID = $1", id)
	u := &domain.Format{}
	err := row.Scan(&u.ID, &u.Title)
	if err != nil {
		return nil, err
	}

	return u, nil
}
