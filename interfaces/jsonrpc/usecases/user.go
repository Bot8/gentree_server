package usecases

import (
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
		useCase usecases.UserUseCase
	}
	ShowUserParams struct{}
	ShowUserResult struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
)

func (h ShowUserHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	u, _ := h.useCase.ShowUser(1)
	return ShowUserResult{
		Id:   u.Id,
		Name: u.Name,
	}, nil
}

func NewShowUser(useCase usecases.UserUseCase) *ShowUserUseCase {
	showUserHandler := &ShowUserHandler{useCase: useCase}
	return &ShowUserUseCase{ShowUserHandler: showUserHandler}
}
