package domain

type UserRepository interface {
	GetAll() ([]User, error)
}

type User struct {
	ID   string
	City City
}
