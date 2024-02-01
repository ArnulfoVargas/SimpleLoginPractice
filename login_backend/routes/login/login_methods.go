package login_routes

import (
	"encoding/json"
	"fullstack_test/dto"
	"net/http"
)

func postLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
  encoder := json.NewEncoder(w)
  loginData := &dto.LoginDto{}

  if err := decoder.Decode(loginData); err != nil {
    response := dto.Response{
      Status: "Error",
    }

    encoder.Encode(response)
    return
  }

  // Search in DB

  // End Search
  response := dto.Response{
    Status: "Ok",
  }
  encoder.Encode(response)
}
