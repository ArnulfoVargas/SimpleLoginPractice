package routes

import (
	login_routes "fullstack_test/routes/login"
	signin_routes "fullstack_test/routes/signin"
	user_routes "fullstack_test/routes/user"

	"github.com/gorilla/mux"
)

func HandleRoutes(r *mux.Router) {
  user_routes.HandleUser(r)
  login_routes.HandleLogin(r)
  signin_routes.HandleSignIn(r)
}
