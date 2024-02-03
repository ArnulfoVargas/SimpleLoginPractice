package middleware

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"fullstack_test/database"
	"fullstack_test/dto"
	"fullstack_test/models"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    encoder := json.NewEncoder(w)

    w.Header().Set("Content-Type", "application/json")
    err := godotenv.Load()
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      encodeMessage(encoder)
      return
    }

    jwtSignature := []byte(os.Getenv("JWT_SIGNATURE"))
    authHeader := r.Header.Get("Authorization")

    if len(authHeader) == 0 {
      w.WriteHeader(http.StatusUnauthorized)
      encoder.Encode(dto.Response{
        Status: "Unauthorized",
      })
      return
    }

    splitBearer := strings.Split(authHeader, " ")
    if len(splitBearer) != 2 {
      w.WriteHeader(http.StatusUnauthorized)
      encoder.Encode(dto.Response{
        Status: "Unauthorized",
      })
      return
    }

    splitToken := strings.Split(splitBearer[1], ".")
    if len(splitToken) != 3 {
      w.WriteHeader(http.StatusUnauthorized)
      encoder.Encode(dto.Response{
        Status: "Unauthorized",
      })
      return
    }

    tk := strings.TrimSpace(splitBearer[1])

    token, err := jwt.Parse(tk, func (tok *jwt.Token) (interface{}, error) {
      if _, ok := tok.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Unexpected Login Error")
      }

      return jwtSignature, nil
    })

    if err != nil {
      w.WriteHeader(http.StatusUnauthorized)
      encoder.Encode(dto.Response{
        Status: "Unauthorized",
      })
      return
    }

    if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
      w.WriteHeader(http.StatusUnauthorized)
      encoder.Encode(dto.Response{
        Status: "Unauthorized",
      })
      return
    } else {
      user := models.User{}
      if err := database.Db.Where("email=?", claims["mail"]).First(&user); err.Error != nil {
        w.WriteHeader(http.StatusUnauthorized)
        encoder.Encode(dto.Response{
          Status: "Unauthorized",
        })
        return
      } 

      next.ServeHTTP(w, r)
    }
  }
}

func encodeMessage(encoder *json.Encoder, errorsFunc ...func(*dto.Response)) {
  res := &dto.Response{}

  for _, errFunc := range errorsFunc{
    errFunc(res)
  }

  encoder.Encode(*res)
}
