package http

import (
	"fmt"
	"net/http"
	"os"
)

func StartHttpServer(host string, port string) {
	address := host + ":" + port
	fmt.Print("Starting server on " + address)

	err := http.ListenAndServe(address, GetRouter())
	if nil != err {
		fmt.Fprintf(os.Stderr, "Unable to start server: %v\n", err)
		os.Exit(1)
	}
}
