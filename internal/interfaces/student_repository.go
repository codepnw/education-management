package interfaces

import "github.com/codepnw/education/internal/entities"

type IStudentRepository interface {
	CreateStudent(student *entities.Student) (string, error)
	GetAllStudents() ([]*entities.Student, error)
	GetStudentByID(id string) (*entities.Student, error)
	UpdateStudentByID(id string, student *entities.Student) error
	DeleteStudent(id string) error
}