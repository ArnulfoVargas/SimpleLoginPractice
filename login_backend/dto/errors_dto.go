package dto

type Errors struct{
  IsValidMail bool `json:"isValidMail,omitempty"`
  IsValidPassword bool `json:"isValidPassword,omitempty"`
  Message string `json:"msg"`
}
