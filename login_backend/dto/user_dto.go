package dto

type UserDto struct {
  Id uint `json:"id"`
  Name string `json:"name" gorm:"type:varchar(50)"`
  Email string `json:"email" gorm:"type:varchar(50)"`
}
