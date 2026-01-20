package domain

type Crew struct {
	ID          string
	Name        string
	CreatedByID string
	CreatedBy   *User
	CreatedAt   string
}
