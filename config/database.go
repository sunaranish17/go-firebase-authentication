package config

import "gorm.io/gorm"

func CreateDatabase() *gorm.DB {

	//Create db instance with gorm
	db, err := gorm.Open("mysql", "go_firebase")
	if err != nil {
		panic("Failed to connect to database!")
	}

	//migrate our model
	db.AutoMigrate()

	return db
}
