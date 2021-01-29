package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SetupDB -> database configuration
func SetupDB() *gorm.DB {
	USER := "root"
	PASS := "admin123"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "go_firebase"

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	log.Print("url", URL)
	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
