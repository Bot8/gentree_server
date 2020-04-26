package rest

import (
	"log"
	"net/http"
)

func StartHttpServer(host string, port string, handler http.Handler) {
	address := host + ":" + port
	log.Println("Starting server on " + address)

	err := http.ListenAndServe(address, handler)
	if nil != err {
		log.Fatalf("Unable to start server: %v\n", err)
	}
}
