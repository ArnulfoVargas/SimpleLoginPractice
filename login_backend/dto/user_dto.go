package dto

type UserDto struct {
  Id uint `json:"id"`
  Username string `json:"username" gorm:"type:varchar(50)"`
  Email string `json:"email" gorm:"type:varchar(50)"`
}
