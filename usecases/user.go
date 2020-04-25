package usecases

import (
	"artarn/gentree/domain/user"
)

type UserInteractor struct {
	UserRepository user.Repository
}

func (interactor *UserInteractor) ShowUser(userId int) (user.User, error) {
	return interactor.UserRepository.FindById(userId)
}

func NewUserInteractor(userRepository user.Repository) *UserInteractor {
	return &UserInteractor{
		UserRepository: userRepository,
	}
}
