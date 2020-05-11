package services

import (
	"artarn/gentree/domain/user"
	"github.com/osamingo/jsonrpc"
)

const (
	ErrorCodeMissingAuthCredentials    jsonrpc.ErrorCode = -4001
	ErrorMessageMissingAuthCredentials string            = "Auth credentials required"

	ErrorCodeInvalidAuthCredentials    jsonrpc.ErrorCode = -4002
	ErrorMessageInvalidAuthCredentials string            = "Invalid auth credentials"
)

type (
	AuthCredentials struct {
		Key int `json:"key"`
	}
	AuthService struct {
		userRepository user.Repository
	}
)

func (service AuthService) GetAuthUser(credentials AuthCredentials) (*user.User, *jsonrpc.Error) {
	if 0 == credentials.Key {
		return nil, ErrMissingAuthCredentials()
	}

	u, err := service.userRepository.FindById(credentials.Key)

	if nil != err {
		return nil, ErrInvalidAuthCredentials()
	}

	return &u, nil
}

func NewAuthService(repository user.Repository) *AuthService {
	return &AuthService{userRepository: repository}
}

func ErrMissingAuthCredentials() *jsonrpc.Error {
	return &jsonrpc.Error{
		Code:    ErrorCodeMissingAuthCredentials,
		Message: ErrorMessageMissingAuthCredentials,
	}
}

func ErrInvalidAuthCredentials() *jsonrpc.Error {
	return &jsonrpc.Error{
		Code:    ErrorCodeInvalidAuthCredentials,
		Message: ErrorMessageInvalidAuthCredentials,
	}
}
