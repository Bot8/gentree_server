package main

import (
	"artarn/gentree/infrastructure"
	"artarn/gentree/infrastructure/database"
	"artarn/gentree/interfaces/database/pg"
	"artarn/gentree/interfaces/http"
	"artarn/gentree/interfaces/http/handlers"
	"artarn/gentree/usecases"
	"log"
)

func main() {
	log.Println("Hello on gentree!")

	config := infrastructure.GetConfig()

	connection := database.GetConnection()

	userRepository := pg.NewPGUserRepository(connection)
	userInteractor := usecases.NewUserInteractor(userRepository)
	userHandler := handlers.NewUserHandler(*userInteractor)

	router := http.GetRouter(*userHandler)

	http.StartHttpServer(config.Server.Host, config.Server.Port, router)
}
