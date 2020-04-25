package handlers

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	_, users := user.GetAllUsers()

	for _, u := range users {
		fmt.Fprintf(w, "user #%d %s", u.Id, u.Name)
	}
}
