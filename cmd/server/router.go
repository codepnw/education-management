package main

import (
	"database/sql"

	"github.com/codepnw/education/internal/handlers"
	"github.com/codepnw/education/internal/repositories"
	"github.com/codepnw/education/internal/usecases"
	"github.com/gin-gonic/gin"
)

type RouterFactory interface {
	StudentsRouter()
	TeachersRouter()
}

type routerFactory struct {
	db *sql.DB
	e  *gin.Engine
}

func NewRouter(db *sql.DB, e *gin.Engine) RouterFactory {
	return &routerFactory{
		db: db,
		e:  e,
	}
}

func (r *routerFactory) StudentsRouter() {
	repo := repositories.NewStudentRepository(r.db)
	usecase := usecases.NewStudentUsecase(repo)
	handler := handlers.NewStudentHandler(usecase)

	router := r.e.Group("/students")

	router.POST("/", handler.CreateStudent)
	router.GET("/", handler.GetAllStudents)
	router.GET("/:id", handler.GetTeacherByID)
	router.PATCH("/:id", handler.UpdateStudentByID)
	router.DELETE("/:id", handler.DeleteStudent)
}

func (r *routerFactory) TeachersRouter() {
	repo := repositories.NewTeacherRepository(r.db)
	usecase := usecases.NewTeacherUsecase(repo)
	handler := handlers.NewTeacherHandler(usecase)

	router := r.e.Group("/teachers")

	router.POST("/", handler.CreateTeacher)
	router.GET("/", handler.GetAllTeachers)
	router.GET("/:id", handler.GetTeacherByID)
	router.PATCH("/:id", handler.UpdateTeacher)
	router.DELETE("/:id", handler.DeleteTeacher)
}
