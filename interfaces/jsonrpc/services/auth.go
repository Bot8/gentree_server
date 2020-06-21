package services

import (
	"artarn/gentree/domain/user"
	"crypto/sha256"
	"fmt"
	"github.com/osamingo/jsonrpc"
)

const (
	ErrorCodeMissingAuthCredentials    jsonrpc.ErrorCode = -4001
	ErrorMessageMissingAuthCredentials string            = "Auth credentials required"

	ErrorCodeInvalidAuthCredentials    jsonrpc.ErrorCode = -4002
	ErrorMessageInvalidAuthCredentials string            = "Invalid auth credentials"

	ErrorCodeUserNotFound    jsonrpc.ErrorCode = -4003
	ErrorMessageUserNotFound string            = "User not found"
)

type (
	AuthCredentials struct {
		AuthToken string `json:"auth_token"`
	}
	AuthService struct {
		userRepository user.Repository
		JWTService
	}
)

func (service AuthService) GetAuthUser(credentials AuthCredentials) (*user.User, *jsonrpc.Error) {
	if "" == credentials.AuthToken {
		return nil, ErrMissingAuthCredentials()
	}

	parsedToken, err := service.JWTService.ParseAuthToken(credentials.AuthToken)

	if nil != err {
		return nil, ErrInvalidAuthCredentials()
	}

	u, err := service.userRepository.FindById(parsedToken.UserId)

	if nil != err {
		return nil, ErrInvalidAuthCredentials()
	}

	return &u, nil
}

func (service AuthService) Login(login string, password string) (*user.User, *jsonrpc.Error) {
	u, err := service.userRepository.FindByLogin(login)

	if nil != err {
		return nil, ErrUserNotFound()
	}

	if false == validatePassword(&u, password) {
		return nil, ErrUserNotFound()
	}

	return &u, nil
}

func validatePassword(user *user.User, password string) bool {
	encodedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	return encodedPassword == user.Password
}

func NewAuthService(repository user.Repository, jwtService JWTService) *AuthService {
	return &AuthService{userRepository: repository, JWTService: jwtService}
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

func ErrUserNotFound() *jsonrpc.Error {
	return &jsonrpc.Error{
		Code:    ErrorCodeUserNotFound,
		Message: ErrorMessageUserNotFound,
	}
}
