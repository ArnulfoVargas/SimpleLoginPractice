package signin_routes

import "github.com/gorilla/mux"

func HandleSignIn(r *mux.Router) {
  r.HandleFunc("/signin", postSignIn)
}
