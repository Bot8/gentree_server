package usecases

import (
	"artarn/gentree/interfaces/jsonrpc/services"
	"artarn/gentree/usecases"
	"context"
	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
)

type (
	ShowUserUseCase struct {
		ShowUserHandler *ShowUserHandler
		ShowUserParams  ShowUserParams
		ShowUserResult  ShowUserResult
	}
	ShowUserHandler struct {
		useCase     usecases.UserUseCase
		authService services.AuthService
	}
	ShowUserParams struct {
		services.AuthCredentials `json:"auth"`
	}
	ShowUserResult struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
)

func (h ShowUserHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var p ShowUserParams
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

func NewShowUser(useCase usecases.UserUseCase, authService services.AuthService) *ShowUserUseCase {
	showUserHandler := &ShowUserHandler{useCase: useCase, authService: authService}
	return &ShowUserUseCase{ShowUserHandler: showUserHandler}
}
