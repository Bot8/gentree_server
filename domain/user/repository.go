package user

type Repository interface {
	Store(user User)
	FindById(id int) User
}
