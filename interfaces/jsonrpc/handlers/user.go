package handlers

import (
	"artarn/gentree/usecases"
	"context"
	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
)

type (
	ShowUser struct {
		ShowUserHandler *ShowUserHandler
		ShowUserParams  ShowUserParams
		ShowUserResult  ShowUserResult
	}
	ShowUserHandler struct {
		interactor usecases.UserInteractor
	}
	ShowUserParams struct{}
	ShowUserResult struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
)

func (h ShowUserHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	u, _ := h.interactor.ShowUser(1)
	return ShowUserResult{
		Id:   u.Id,
		Name: u.Name,
	}, nil
}

func NewShowUser(interactor usecases.UserInteractor) *ShowUser {
	showUserHandler := &ShowUserHandler{interactor: interactor}
	return &ShowUser{ShowUserHandler: showUserHandler}
}
