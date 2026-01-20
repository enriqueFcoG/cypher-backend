package domain

type History struct {
	ID       string
	IDUser   string
	User     *User
	IDFormat string
	Format   *Format
}
