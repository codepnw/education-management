package interfaces

import "github.com/codepnw/education/internal/entities"

type CourseRepository interface {
	CreateCourse(req *entities.Course) (string, error)
	GetAllCourses() ([]*entities.Course, error)
	GetCourseByID(id string) (*entities.Course, error)
	UpdateCourse(id string, req *entities.Course) error
	DeleteCourse(id string) error
}