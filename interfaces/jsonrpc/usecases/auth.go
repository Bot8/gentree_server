package usecases

import (
	"artarn/gentree/interfaces/jsonrpc/services"
	"context"
	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
)

type (
	LoginUseCase struct {
		LoginHandler *LoginHandler
		LoginParams  LoginParams
		LoginResult  LoginResult
	}
	LoginHandler struct {
		authService services.AuthService
		jwtService  services.JWTService
	}
	LoginParams struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	LoginResult struct {
		Id         int                 `json:"id"`
		Name       string              `json:"name"`
		AuthTokens services.AuthTokens `json:"auth_tokens"`
	}
)

func (h LoginHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var p LoginParams
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	u, err := h.authService.Login(p.Login, p.Password)

	if nil != err {
		return nil, err
	}

	authTokens := services.AuthTokens{
		AuthToken: h.jwtService.GetAuthToken(u),
	}

	return LoginResult{
		Id:         u.Id,
		Name:       u.Name,
		AuthTokens: authTokens,
	}, nil
}

func NewLoginUseCase(authService *services.AuthService, jwtService *services.JWTService) *LoginUseCase {
	return &LoginUseCase{
		LoginHandler: &LoginHandler{authService: *authService, jwtService: *jwtService},
	}
}