package main

import "artarn/gentree/domain/user"

type UserInteractor struct {
	UserRepository user.Repository
}

func (interactor *UserInteractor) ShowUser(userId int) user.User {
	user := interactor.UserRepository.FindById(userId)

	return user
}
