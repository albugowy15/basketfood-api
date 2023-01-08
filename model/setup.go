package model

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDB() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	dsn := "host=" + os.Getenv("DB_HOST") + " user=postgres password=" + os.Getenv("DB_PASSWORD")+ " dbname=postgres port=5432"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&Ecommerce{}, &Product{}, &Staff{}, &Customer{}, &Report{}, &Discout{}, &EcommerceAccount{}, &Order{}, &ProductOrder{})

	DB = db
}

