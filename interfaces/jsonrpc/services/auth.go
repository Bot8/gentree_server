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
		Key string `json:"key"`
	}
	AuthTokens struct {
		AuthToken      string `json:"auth"`
		RefreshToken   string `json:"refresh"`
		SignerToken    string `json:"signer_token"`
		EncryptedToken string `json:"encrypted_token"`
	}
	AuthService struct {
		userRepository user.Repository
		JWTService
	}
)

func (service AuthService) GetAuthUser(credentials AuthCredentials) (*user.User, *jsonrpc.Error) {
	if "" == credentials.Key {
		return nil, ErrMissingAuthCredentials()
	}

	parsedToken, err := service.JWTService.ParseAuthToken(credentials.Key)

	if nil != err {
		return nil, ErrInvalidAuthCredentials()
	}

	u, err := service.userRepository.FindById(parsedToken.UserId)

	if nil != err {
		return nil, ErrInvalidAuthCredentials()
	}

	return &u, nil
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
