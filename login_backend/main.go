package main

import (
	"fmt"
	"fullstack_test/routes"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	muxS := mux.NewRouter()

  routes.HandleRoutes(muxS)

  // Set up Cors
  handler := cors.AllowAll().Handler(muxS)

  fmt.Println("Server running...")
  panic(http.ListenAndServe(generateHost(),handler))
}

func generateHost() string {
  godotenv.Load()
  sb := strings.Builder{}

  sb.WriteString(os.Getenv("HOST"))
  sb.WriteByte(':')
  sb.WriteString(os.Getenv("PORT"))

  return sb.String()
}
