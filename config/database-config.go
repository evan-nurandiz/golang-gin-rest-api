package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Setup Database Connection
func SetupDatabaseConnection() *gorm.DB {
	dsn := "host=localhost user=evan password=bangroni dbname=golang_api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
		panic("Failed to connect to db")
	}

	return db
}

//Close Database Connection
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to Close Connection From Database")
	}

	dbSQL.Close()
}
