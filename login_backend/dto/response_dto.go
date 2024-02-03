package dto

type Response struct {
	Status string `json:"status"`
	Errors `json:"errors"`
  User UserDto `json:"user,omitempty"`
  Token string `json:"token"`
}

func NewResponse() *Response {
  return &Response{
    Errors: Errors{
      IsValidMail: true,
      IsValidPassword: true,
    },
  }
}
