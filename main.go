package main

import (
	"artarn/gentree/infrastructure"
	"artarn/gentree/infrastructure/database"
	"artarn/gentree/interfaces/database/pg"
	"artarn/gentree/interfaces/jsonrpc"
	jsonrpcUsecases "artarn/gentree/interfaces/jsonrpc/usecases"
	"artarn/gentree/usecases"

	"log"
)

func main() {
	log.Println("Hello on gentree!")

	config := infrastructure.GetConfig()

	connection := database.GetConnection()

	userRepository := pg.NewPGUserRepository(connection)

	userUseCase := usecases.NewUserUseCase(userRepository)

	showUser := jsonrpcUsecases.NewShowUser(*userUseCase)
	methodRepository := jsonrpc.GetNewMethodRepository(*showUser)

	jsonrpc.StartJSONRPCServer(config.JsonRPRCServer.Host, config.JsonRPRCServer.Port, methodRepository)
}
