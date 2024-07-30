package main

import (
	"log"
	"os"

	"github.com/codepnw/education/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed loading .env file: %v", err)
	}

	if err := database.ConnectDatabase(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.Run(":" + os.Getenv("PORT"))
}
