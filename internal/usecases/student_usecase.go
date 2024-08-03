package usecases

import (
	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/interfaces"
)

type StudentUsecase struct {
	repo interfaces.IStudentRepository
}

func NewStudentUsecase(repo interfaces.IStudentRepository) *StudentUsecase {
	return &StudentUsecase{repo: repo}
}

func (su *StudentUsecase) CreateStudent(student *entities.Student) (string, error) {
	return su.repo.CreateStudent(student)
}

func (su *StudentUsecase) GetAllStudents() ([]*entities.Student, error) {
	return su.repo.GetAllStudents()
}

func (su *StudentUsecase) GetStudentByID(id string) (*entities.Student, error) {
	return su.repo.GetStudentByID(id)
}

func (su *StudentUsecase) UpdateStudentByID(id string, student *entities.Student) error {
	return su.repo.UpdateStudentByID(id, student)
}

func (su *StudentUsecase) DeleteStudent(id string) error {
	return su.repo.DeleteStudent(id)
}