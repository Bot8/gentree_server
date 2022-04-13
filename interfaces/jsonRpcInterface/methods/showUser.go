package methods

import (
	"artarn/gentree/interfaces/jsonRpcInterface/services"
	"context"
	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
)

type (
	ShowUserMethod struct {
		ShowUserHandler *ShowUserHandler
		ShowUserParams  ShowUserParams
		ShowUserResult  ShowUserResult
	}
	ShowUserHandler struct {
		authService services.AuthService
	}
	ShowUserParams struct {
		AuthParams
	}
	ShowUserResult struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
)

func CreateShowUserMethod(authService *services.AuthService) *ShowUserMethod {
	return &ShowUserMethod{
		ShowUserHandler: &ShowUserHandler{authService: *authService},
	}
}

func (h ShowUserHandler) ServeJSONRPC(_ context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var p AuthParams
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	u, err := h.authService.GetAuthUser(p.AuthCredentials)
	if nil != err {
		return nil, err
	}

	return ShowUserResult{
		Id:   u.Id,
		Name: u.Name,
	}, nil
}
