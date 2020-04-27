package usecases

import (
	"artarn/gentree/domain/user"
)

type UserUseCase struct {
	UserRepository user.Repository
}

func (useCase *UserUseCase) ShowUser(userId int) (user.User, error) {
	return useCase.UserRepository.FindById(userId)
}

func NewUserUseCase(userRepository user.Repository) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}
