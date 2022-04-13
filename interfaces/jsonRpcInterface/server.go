package jsonRpcInterface

import (
	"github.com/osamingo/jsonrpc"
	"goji.io"
	"goji.io/pat"
	"log"
	"net/http"
)

func StartJSONRPCServer(host string, port string, handler *jsonrpc.MethodRepository) {
	address := host + ":" + port
	log.Println("Starting JSON RPC server on " + address)

	mux := goji.NewMux()

	mux.Handle(pat.Post("/jrpc"), handler)
	mux.HandleFunc(pat.Post("/jrpc/debug"), handler.ServeDebug)

	err := http.ListenAndServe(address, mux)
	if nil != err {
		log.Fatalf("Unable to start server: %v\n", err)
	}
}
