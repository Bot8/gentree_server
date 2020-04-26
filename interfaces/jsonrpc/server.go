package jsonrpc

import (
	"github.com/osamingo/jsonrpc"
	"log"
	"net/http"
)

func StartJSONRPCServer(host string, port string, handler *jsonrpc.MethodRepository) {
	address := host + ":" + port
	log.Println("Starting JSON RPC server on " + address)

	http.Handle("/jrpc", handler)
	http.HandleFunc("/jrpc/debug", handler.ServeDebug)

	err := http.ListenAndServe(address, http.DefaultServeMux)
	if nil != err {
		log.Fatalf("Unable to start server: %v\n", err)
	}
}
