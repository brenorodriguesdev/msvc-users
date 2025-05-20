package main

import (
	"fmt"
	"log"
	"os"

	database "msvc-users/main/config"
	"msvc-users/main/routes"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using default port")
	}

	database.DB = database.InitDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	routes.UserRoutes(r)

	err = r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
