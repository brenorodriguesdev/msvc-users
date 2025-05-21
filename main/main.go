package main

import (
	"fmt"
	"log"
	"os"

	database "msvc-users/main/config"
	"msvc-users/main/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Só tenta carregar .env se a variável ENV for "local" ou não estiver definida
	env := os.Getenv("ENV")
	if env == "" || env == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: .env file not found, using default port")
		}
	}

	database.DB = database.InitDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	routes.UserRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
