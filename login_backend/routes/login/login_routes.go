package login_routes

import (
	"fullstack_test/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleLogin(r *mux.Router) {
  r.HandleFunc("/login", postLogin).Methods(http.MethodPost)
  r.HandleFunc("/jwt", middleware.ValidateJWT(loginByJWT))
}
