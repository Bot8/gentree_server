package handlers

import (
	"artarn/gentree/usecases"
	"fmt"
	"net/http"
)

type UserHandler struct {
	interactor usecases.UserInteractor
}

func NewUserHandler(interactor usecases.UserInteractor) *UserHandler {
	return &UserHandler{interactor: interactor}
}

func (h *UserHandler) ShowUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		u, _ := h.interactor.ShowUser(1)

		_, _ = fmt.Fprintf(w, "user #%d %s", u.Id, u.Name)
	}
}
