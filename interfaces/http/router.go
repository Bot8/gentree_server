package http

import (
	"net/http"

	"github.com/gorilla/mux"

	"artarn/gentree/interfaces/http/handlers"
)

func GetRouter(userHandler handlers.UserHandler) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", userHandler.ShowUser())

	return router
}
