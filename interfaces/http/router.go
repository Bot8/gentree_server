package http

import (
	"net/http"

	"github.com/gorilla/mux"

	"artarn/gentree/interfaces/http/handlers"
)

func GetRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.IndexHandler)

	return router
}
