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
	ClassroomRouter()
	CourseRouter()
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

func (r *routerFactory) ClassroomRouter() {
	repo := repositories.NewClassroomRepository(r.db)
	usecase := usecases.NewClassroomUsecase(repo)
	handler := handlers.NewClassroomHandler(usecase)

	router := r.e.Group("/classroom")

	router.POST("/", handler.CreateClassroom)
	router.GET("/", handler.GetAllClassroom)
	router.GET("/:id", handler.GetClassroomByID)
	router.PATCH("/:id", handler.UpdateClassroom)
	router.DELETE("/:id", handler.DeleteClassroom)
}

func (r *routerFactory) CourseRouter() {
	repo := repositories.NewCourseRepository(r.db)
	usecase := usecases.NewCourseUsecase(repo)
	handler := handlers.NewCourseHandler(usecase)

	router := r.e.Group("/courses")

	router.POST("/", handler.CreateCourse)
	router.GET("/", handler.GetAllCourses)
	router.GET("/:id", handler.GetCourseByID)
	router.PATCH("/:id", handler.UpdateCourse)
	router.DELETE("/:id", handler.DeleteCourse)
}