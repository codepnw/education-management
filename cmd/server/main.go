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

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()
	router := NewRouter(db, r)
	
	router.StudentsRouter()
	router.TeachersRouter()
	router.ClassroomRouter()
	router.CourseRouter()
	router.EnrollmentRouter()
	router.ScheduleRouter()

	r.Run(":" + os.Getenv("PORT"))
}
