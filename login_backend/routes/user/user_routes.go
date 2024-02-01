package user_routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HandleUser(r *mux.Router) {
  prefix := "/api/test"

  r.HandleFunc(prefix, getUser).Methods(http.MethodGet) 
  r.HandleFunc(prefix, postUser).Methods(http.MethodPost) 
}
