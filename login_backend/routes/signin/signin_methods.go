package signin_routes

import (
	"encoding/json"
	"fullstack_test/dto"
	"net/http"
)

func postSignIn(w http.ResponseWriter, r *http.Request) {
  decoder := json.NewDecoder(r.Body)
  signIn := dto.SignInDto{}

  if err := decoder.Decode(&signIn); err != nil {

  }
}
