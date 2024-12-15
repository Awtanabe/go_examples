package main

import (
	db "practice_todo/db"
	"practice_todo/models"
)


func main() {
	dbConn := db.NewDB()

	dbConn.AutoMigrate(&models.User{}, &models.Todo{})
}