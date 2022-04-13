package main

import (
	"artarn/gentree/dataAccess/pg"
	"artarn/gentree/infrastructure"
	"artarn/gentree/infrastructure/database"
	"artarn/gentree/interfaces/jsonRpcInterface"
	"artarn/gentree/interfaces/jsonRpcInterface/methods"
	"artarn/gentree/interfaces/jsonRpcInterface/services"

	"log"
)

func main() {
	log.Println("Hello on gentree!")

	config := infrastructure.GetConfig()

	connection := database.GetConnection()
	userRepository := pg.CreateUserRepository(connection)

	jwtService := services.CreateJWTService(config.Encryption.JWTSecret)
	authService := services.CreateAuthService(userRepository, *jwtService)

	showUserMethod := methods.CreateShowUserMethod(authService)
	loginMethod := methods.CreateLoginMethod(authService, jwtService)

	methodRepository := jsonRpcInterface.GetNewMethodRepository(*showUserMethod, *loginMethod)
	jsonRpcInterface.StartJSONRPCServer(config.JsonRPRCServer.Host, config.JsonRPRCServer.Port, methodRepository)
}
