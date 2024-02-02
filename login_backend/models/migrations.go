package models

import "fullstack_test/database"

func DbMigrate() {
	database.Db.AutoMigrate(&User{})
}
