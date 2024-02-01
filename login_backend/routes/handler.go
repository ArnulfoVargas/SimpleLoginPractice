package routes

import (
	user_routes "fullstack_test/routes/user"

	"github.com/gorilla/mux"
)

func HandleRoutes(r *mux.Router) {
  user_routes.HandleUser(r)
}
