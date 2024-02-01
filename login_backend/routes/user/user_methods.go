package user_routes

import (
	"encoding/json"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

  encoder.Encode(map[string]any{
    "status": http.StatusOK,
    "message": "Success on fetch",
  })
}

func postUser(w http.ResponseWriter, r *http.Request) {
  decoder := json.NewDecoder(r.Body)
  encoder := json.NewEncoder(w)
  data := make(map[string]any)

  if err := decoder.Decode(&data); err != nil {
    encoder.Encode(map[string]any{
      "status": http.StatusConflict,
      "message": "Error on post",
    })
    return
  }

  encoder.Encode(map[string]any{
    "status": http.StatusOK,
    "message": "Success on Post",
  })
}
