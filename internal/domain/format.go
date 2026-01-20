package domain

type Format struct {
	ID          string
	Title       string
	IsPublic    bool
	Likes       int
	CreatedByID string
	CreatedBy   *User
	CreatedAt   string
}
