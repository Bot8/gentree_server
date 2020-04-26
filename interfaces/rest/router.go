package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"artarn/gentree/interfaces/rest/handlers"
)

func GetNewRouter(userHandler handlers.UserHandler) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", userHandler.ShowUser())

	return router
}
