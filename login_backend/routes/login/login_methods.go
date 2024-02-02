package login_routes

import (
	"encoding/json"
	"fullstack_test/database"
	"fullstack_test/dto"
	"fullstack_test/models"
	"net/http"
  "golang.org/x/crypto/bcrypt"
)

func postLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
  encoder := json.NewEncoder(w)
  deps := &responseDeps{
    w: w,
    encoder: encoder,
  }
  loginData := &dto.LoginDto{}

  if err := decoder.Decode(loginData); err != nil {
    generateResponse(deps, http.StatusBadRequest, setStatusType("Error"))
    return
  }

  // Search in DB
  user := models.User{}

  if database.Db.Where("email=?", user.Email).Limit(1).Find(&user).RowsAffected <= 0 {
    generateResponse(deps, http.StatusNotFound, setInvalidMail, userNotFound)
    return
  }
  // End Search

  if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
    generateResponse(deps, http.StatusNotFound, setInvalidPassword, setInvalidMail)
    return
  }

  response := dto.Response{
    Status: "Ok",
    User: dto.UserDto{
      Id: user.Id,
      Name: user.Name,
      Email: user.Email,
    },
  }
  encoder.Encode(response)
}

type responseDeps struct {
  encoder *json.Encoder
  w http.ResponseWriter
}

func generateResponse(deps *responseDeps, statusCode int, opts ...func(*dto.Response)) {
  response := dto.NewResponse()

  for _, opt := range opts {
    opt(response)
  }

  deps.w.WriteHeader(statusCode)
  deps.encoder.Encode(*response)
}

func setStatusType(status string) func(*dto.Response) {
  return func(r *dto.Response) {
    r.Status = status
  }
}

func setInvalidMail(r *dto.Response) {
  r.Status = "Error"
  r.Errors.IsValidMail = false
  r.Errors.Message = r.Errors.Message + "Invalid E-Mail."
}

func setInvalidPassword(r *dto.Response) {
  r.Status = "Error"
  r.Errors.IsValidPassword = false
  r.Errors.Message = r.Errors.Message + "Invalid Password."
}

func userNotFound(r *dto.Response) {
  r.Status = "User Not Found"
}
