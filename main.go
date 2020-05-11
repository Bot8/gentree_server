package main

import (
	"artarn/gentree/infrastructure"
	"artarn/gentree/infrastructure/database"
	"artarn/gentree/interfaces/database/pg"
	"artarn/gentree/interfaces/jsonrpc"
	"artarn/gentree/interfaces/jsonrpc/services"
	jsonRPCUseCases "artarn/gentree/interfaces/jsonrpc/usecases"
	domainUseCases "artarn/gentree/usecases"

	"log"
)

func main() {
	log.Println("Hello on gentree!")

	config := infrastructure.GetConfig()

	connection := database.GetConnection()

	userRepository := pg.NewPGUserRepository(connection)
	userUseCase := domainUseCases.NewUserUseCase(userRepository)

	authService := services.NewAuthService(userRepository)

	showUser := jsonRPCUseCases.NewShowUser(*userUseCase, *authService)
	methodRepository := jsonrpc.GetNewMethodRepository(*showUser)

	jsonrpc.StartJSONRPCServer(config.JsonRPRCServer.Host, config.JsonRPRCServer.Port, methodRepository)
}
