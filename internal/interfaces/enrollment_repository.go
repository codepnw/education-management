package interfaces

import "github.com/codepnw/education/internal/entities"

type EnrollmentRepository interface {
	CreateEnroll(req *entities.Enrollment) (string, error)
	GetAllEnroll() ([]*entities.Enrollment, error)
	GetEnrollByID(id string) (*entities.Enrollment, error)
	UpdateEnroll(id string, req *entities.Enrollment) error
	DeleteEnroll(id string) error
}