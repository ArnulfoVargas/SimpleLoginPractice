package dto

type Response struct {
  Status string `json:"status"`
  Errors `json:"errors,omitempty"`
  UserId uint `json:"userId,omitempty"`
  Username string `json:"username,omitempty"`
}
