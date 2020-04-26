package jsonrpc

import (
	"artarn/gentree/interfaces/jsonrpc/handlers"
	"github.com/osamingo/jsonrpc"
	"log"
)

func GetNewMethodRepository(showUser handlers.ShowUser) *jsonrpc.MethodRepository {
	jsonRPCServer := jsonrpc.NewMethodRepository()

	err := jsonRPCServer.RegisterMethod("User.ShowInfo", showUser.ShowUserHandler, showUser.ShowUserParams, showUser.ShowUserResult)

	if nil != err {
		log.Fatalln(err)
	}

	return jsonRPCServer
}
