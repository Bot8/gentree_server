package main

import (
	"artarn/gentree/infrastructure"
	"artarn/gentree/infrastructure/database"
	"artarn/gentree/interfaces/database/pg"
	"artarn/gentree/interfaces/jsonrpc"
	handlersJsonrpc "artarn/gentree/interfaces/jsonrpc/handlers"
	"artarn/gentree/interfaces/rest"
	"artarn/gentree/interfaces/rest/handlers"
	"artarn/gentree/usecases"

	"log"
)

func main() {
	log.Println("Hello on gentree!")

	config := infrastructure.GetConfig()

	connection := database.GetConnection()

	userRepository := pg.NewPGUserRepository(connection)
	userInteractor := usecases.NewUserInteractor(userRepository)

	if config.RestServer.Enabled {
		userHandler := handlers.NewUserHandler(*userInteractor)
		router := rest.GetNewRouter(*userHandler)
		rest.StartHttpServer(config.RestServer.Host, config.RestServer.Port, router)
	}

	if config.JsonRPRCServer.Enabled {
		showUser := handlersJsonrpc.NewShowUser(*userInteractor)
		methodRepository := jsonrpc.GetNewMethodRepository(*showUser)
		jsonrpc.StartJSONRPCServer(config.JsonRPRCServer.Host, config.JsonRPRCServer.Port, methodRepository)
	}
}
