package interfaces

import "github.com/codepnw/education/internal/entities"

type ClassroomRepository interface {
	CreateClassroom(req *entities.Classroom) (string, error)
	GetAllClassroom() ([]*entities.Classroom, error)
	GetClassroomByID(id string) (*entities.Classroom, error)
	UpdateClassroom(id string, req *entities.Classroom) error
	DeleteClassroom(id string) error
}
