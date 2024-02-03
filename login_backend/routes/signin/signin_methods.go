package signin_routes

import (
	"encoding/json"
	"fullstack_test/database"
	"fullstack_test/dto"
	jwt_generator "fullstack_test/jwt"
	"fullstack_test/models"
	"fullstack_test/validations"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func postSignIn(w http.ResponseWriter, r *http.Request) {
  decoder := json.NewDecoder(r.Body)
  encoder := json.NewEncoder(w)
  signIn := dto.SignInDto{}
  deps := &deps{
    w: w,
    encoder: encoder,
  }

  if err := decoder.Decode(&signIn); err != nil {
    generateResponse(deps, http.StatusNotAcceptable, setStatusAsError)
    return
  }

  // Search in DB
  user := models.User{}
  if database.Db.Where("email=?", signIn.Email).Find(&user).Limit(1).RowsAffected > 0 {
    generateResponse(deps, http.StatusBadRequest, setInvalidMail("Email already in use"))
    return
  }
  // If not exist in db

  match := validations.Mail_Regex.FindStringSubmatch(signIn.Email)
  isValidPassword := validations.ValidatePassword(signIn.Password)

  switch {
  case match == nil && !isValidPassword:
    generateResponse(deps, http.StatusBadRequest, setInvalidMail(""), setInvalidPassword)
    return
  case match == nil:
    generateResponse(deps, http.StatusBadRequest, setInvalidMail(""))
    return
  case !isValidPassword:
    generateResponse(deps, http.StatusBadRequest, setInvalidPassword)
    return
  }

  var cost = 8
  crypPassword, err := bcrypt.GenerateFromPassword([]byte(signIn.Password), cost)
  if err != nil {
    generateResponse(deps, http.StatusConflict, setStatusAsError)
    return
  }

  newUser := models.User{
    Username: signIn.Username,
    Password: string(crypPassword),
    Email: signIn.Email,
  }
  database.Db.Save(&newUser)


  res := dto.Response{
    Status: "Created",
  }

  w.WriteHeader(http.StatusCreated)
  if database.Db.Where("email=?", signIn.Email).Find(&user).Limit(1).RowsAffected < 1 {
    encoder.Encode(res)
    return
  }

  userDto := dto.UserDto{
    Id: user.Id,
    Username: user.Username,
    Email: user.Email,
  }

  res.User = userDto

  token, err := jwt_generator.GenerateJWT(user)
  if err != nil {
    encoder.Encode(res)
    return
  }
  res.Token = token

  encoder.Encode(res)
}

type deps struct {
  w http.ResponseWriter
  encoder *json.Encoder
}

func generateResponse(deps *deps, statusCode int, opts ...func(*dto.Response)) {
  res := &dto.Response{}

  for _, opt := range opts {
    opt(res)
  }

  deps.w.WriteHeader(statusCode)
  deps.encoder.Encode(res)
}

func setStatusAsError(r *dto.Response) {
  r.Status = "Error"
}

func setInvalidPassword(r *dto.Response) {
  r.Errors.IsValidPassword = false
  r.Errors.Message = r.Errors.Message + "Invalid Password."
}

func setInvalidMail(message string) func(*dto.Response) {
  var msg string
  if len(message) > 2 {
    msg = message
  } else {
    msg = "Invalid Mail."
  }
  return func (r *dto.Response) {
    r.Errors.IsValidMail = false
    r.Errors.Message = r.Errors.Message +  msg
  }
}

