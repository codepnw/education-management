package usecases

import (
	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/interfaces"
)

type ClassroomUsecase struct {
	repo interfaces.ClassroomRepository
}

func NewClassroomUsecase(repo interfaces.ClassroomRepository) *ClassroomUsecase {
	return &ClassroomUsecase{repo: repo}
}

func (cu *ClassroomUsecase) CreateClassroom(req *entities.Classroom) (string, error) {
	return cu.repo.CreateClassroom(req)
}

func (cu *ClassroomUsecase) GetAllClassroom() ([]*entities.Classroom, error) {
	return cu.repo.GetAllClassroom()
}

func (cu *ClassroomUsecase) GetClassroomByID(id string) (*entities.Classroom, error) {
	return cu.repo.GetClassroomByID(id)
}

func (cu *ClassroomUsecase) UpdateClassroom(id string, req *entities.Classroom) error {
	return cu.repo.UpdateClassroom(id, req)
}

func (cu *ClassroomUsecase) DeleteClassroom(id string) error {
	return cu.repo.DeleteClassroom(id)
}

