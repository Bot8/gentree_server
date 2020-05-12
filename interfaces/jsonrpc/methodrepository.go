package jsonrpc

import (
	"artarn/gentree/interfaces/jsonrpc/usecases"
	"github.com/osamingo/jsonrpc"
)

func GetNewMethodRepository(showUser usecases.ShowUserUseCase, login usecases.LoginUseCase) *jsonrpc.MethodRepository {
	jsonRPCServer := jsonrpc.NewMethodRepository()

	jsonRPCServer.RegisterMethod("User.ShowInfo", showUser.ShowUserHandler, showUser.ShowUserParams, showUser.ShowUserResult)
	jsonRPCServer.RegisterMethod("User.Login", login.LoginHandler, login.LoginParams, login.LoginResult)

	//if nil != err {
	//	log.Fatalln(err)
	//}

	return jsonRPCServer
}
