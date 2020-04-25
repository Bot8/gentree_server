package main

import (
	"artarn/gentree/infrastructure"
	"artarn/gentree/interfaces/http"
	"fmt"
)

func main() {
	fmt.Println("Hello on gentree!")

	serverConfig := infrastructure.GetConfig().Server

	http.StartHttpServer(serverConfig.Host, serverConfig.Port)
}
