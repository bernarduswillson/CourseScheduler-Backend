package initializers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	DB, err = gorm.Open(mysql.Open(dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbTable+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	log.Println("database connection successful")
}
