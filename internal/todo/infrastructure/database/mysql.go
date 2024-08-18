package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPasswored := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	dsn := dbUser + ":" + dbPasswored + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
