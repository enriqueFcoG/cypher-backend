package domain

type Mod struct {
	ID          string
	Title       string
	Description string
	CreatedAt   string
	CreatedByID string
	CreatedBy   *User
}
