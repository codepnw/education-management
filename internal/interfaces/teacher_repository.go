package interfaces

import "github.com/codepnw/education/internal/entities"

type TeacherRepository interface {
	CreateTeacher(teacher *entities.Teacher) (string, error)
	GetAllTeachers() ([]*entities.Teacher, error)
	GetTeacherByID(id string) (*entities.Teacher, error)
	UpdateTeacher(id string, teacher *entities.Teacher) error
	DeleteTeacher(id string) error
}
