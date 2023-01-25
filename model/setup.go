package model

import (
	"fmt"
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
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD")+ " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Migrating db")

	db.AutoMigrate(&Ecommerce{}, &Product{}, &Staff{}, &Customer{}, &Report{}, &Discout{}, &EcommerceAccount{}, &Order{}, &ProductOrder{})

	DB = db
}

