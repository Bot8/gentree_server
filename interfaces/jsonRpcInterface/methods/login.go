package methods

import (
	"artarn/gentree/interfaces/jsonRpcInterface/services"
	"context"
	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
)

type (
	AuthParams struct {
		services.AuthCredentials `json:"auth"`
	}
	LoginMethod struct {
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
		Id              int                      `json:"id"`
		Name            string                   `json:"name"`
		AuthCredentials services.AuthCredentials `json:"auth"`
	}
)

func CreateLoginMethod(authService *services.AuthService, jwtService *services.JWTService) *LoginMethod {
	return &LoginMethod{
		LoginHandler: &LoginHandler{authService: *authService, jwtService: *jwtService},
	}
}

func (h LoginHandler) ServeJSONRPC(_ context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var p LoginParams
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	u, err := h.authService.Login(p.Login, p.Password)

	if nil != err {
		return nil, err
	}

	authCredentials := services.AuthCredentials{
		AuthToken: h.jwtService.GetAuthToken(u),
	}

	return LoginResult{
		Id:              u.Id,
		Name:            u.Name,
		AuthCredentials: authCredentials,
	}, nil
}
