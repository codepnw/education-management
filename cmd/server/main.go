package main

import (
	"log"
	"os"

	"github.com/codepnw/education/internal/handlers"
	"github.com/codepnw/education/internal/repositories"
	"github.com/codepnw/education/internal/usecases"
	"github.com/codepnw/education/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed loading .env file: %v", err)
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stdRepo := repositories.NewStudentRepository(db)
	stdUsecase := usecases.NewStudentUsecase(stdRepo)
	stdHandler := handlers.NewStudentHandler(stdUsecase)

	r := gin.Default()

	r.POST("/students", stdHandler.CreateStudent)
	r.GET("/students", stdHandler.GetAllStudents)
	r.PATCH("/students/:id", stdHandler.UpdateStudentByID)
	r.DELETE("/students/:id", stdHandler.DeleteStudent)

	r.Run(":" + os.Getenv("PORT"))
}
