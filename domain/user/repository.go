package user

type Repository interface {
	Store(user User) error
	FindById(id int) (User, error)
}
