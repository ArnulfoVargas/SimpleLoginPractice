package login_routes

import "github.com/gorilla/mux"

func HandleLogin(r *mux.Router) {
  r.HandleFunc("/login", postLogin)
}
