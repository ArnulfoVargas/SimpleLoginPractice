package jwt_generator

import (
	"fmt"
	"fullstack_test/models"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateJWT(user models.User) (string, error){
  err := godotenv.Load()
  if err != nil {
    panic(err)
  }
  
  key := []byte(os.Getenv("JWT_SIGNATURE"))

  generatedtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "mail"  : user.Email,
    "username"  : user.Username,
    "id"    : user.Id,
    "iat"   : time.Now().Unix(),
    "exp"   : time.Now().Add(time.Minute * 5).Unix(),
  })

  jsonToken, err := generatedtoken.SignedString(key) 
  fmt.Println(err)
  return jsonToken, err
}
