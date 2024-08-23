package controllers

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"log"

)
var Db *gorm.DB
func ConnectDatabase() {
	// Connect to the database
	dsn := "root:050220@tcp(127.0.0.1:3306)/mydb?parseTime=true&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := database.AutoMigrate(&StudentScore{}, &Course{}, &Student{}); err != nil {
		log.Fatal(err)
	}
	Db = database
}