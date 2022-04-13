package jsonRpcInterface

import (
	"artarn/gentree/interfaces/jsonRpcInterface/methods"
	"github.com/osamingo/jsonrpc"
)

func GetNewMethodRepository(showUserMethod methods.ShowUserMethod, loginMethod methods.LoginMethod) *jsonrpc.MethodRepository {
	jsonRPCServer := jsonrpc.NewMethodRepository()

	jsonRPCServer.RegisterMethod(
		"User.ShowInfo",
		showUserMethod.ShowUserHandler,
		showUserMethod.ShowUserParams,
		showUserMethod.ShowUserResult,
	)

	jsonRPCServer.RegisterMethod(
		"User.Login",
		loginMethod.LoginHandler,
		loginMethod.LoginParams,
		loginMethod.LoginResult,
	)

	return jsonRPCServer
}
