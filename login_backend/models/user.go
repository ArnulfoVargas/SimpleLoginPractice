package models

type User struct{
  Id uint `json:"id"`
  Name string `json:"name" gorm:"type:varchar(50)"`
  Email string `json:"email" gorm:"type:varchar(50)"`
  Password string `json:"password"`
}
