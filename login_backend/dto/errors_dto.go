package dto

type Errors struct{
  IsValidUserName bool `json:"isValidUserName,omitempty"`
  IsValidMail bool `json:"isValidMail,omitempty"`
  IsValidPassword bool `json:"isValidPassword,omitempty"`
}
