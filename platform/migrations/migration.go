package main

import (
	"auth/app/models"
	"auth/platform/database"
)

func main() {
	db, err := database.MysqlConnection()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
}
