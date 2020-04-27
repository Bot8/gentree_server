package jsonrpc

import (
	"artarn/gentree/interfaces/jsonrpc/usecases"
	"github.com/osamingo/jsonrpc"
	"log"
)

func GetNewMethodRepository(showUser usecases.ShowUserUseCase) *jsonrpc.MethodRepository {
	jsonRPCServer := jsonrpc.NewMethodRepository()

	err := jsonRPCServer.RegisterMethod("User.ShowInfo", showUser.ShowUserHandler, showUser.ShowUserParams, showUser.ShowUserResult)

	if nil != err {
		log.Fatalln(err)
	}

	return jsonRPCServer
}
