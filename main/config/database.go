package database

import (
	"fmt"
	"log"
	"msvc-users/infra/entities"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate(DB *gorm.DB) error {
	return DB.AutoMigrate(&entities.User{})
}

func InitDB() *gorm.DB {
	env := os.Getenv("ENV")
	if env == "" || env == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: .env file not found, using default env vars")
		}
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, pass, name, port,
	)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = Migrate(DB)
	if err != nil {
		log.Fatal("Error migrating:", err)
	}

	return DB
}
