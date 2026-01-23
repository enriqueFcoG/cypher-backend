package persistence

import (
	"database/sql"
	"enriqueFcoG/cypher/internal/domain"
)

type CrewRepoPG struct {
	DB *sql.DB
}

func (r *CrewRepoPG) Save(u *domain.Crew) error {
	_, error := r.DB.Exec("INSERT INTO CREW (ID, UserName) values($1,$2)", u.ID, u.Name)
	return error
}

func (r *CrewRepoPG) FindByID(id string) (*domain.Crew, error) {
	row := r.DB.QueryRow("SELECT ID, UserName FROM CREW WHERE ID = $1", id)
	u := &domain.Crew{}
	err := row.Scan(&u.ID, &u.Name)
	if err != nil {
		return nil, err
	}

	return u, nil
}
