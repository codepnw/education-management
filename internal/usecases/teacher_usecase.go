package usecases

import (
	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/interfaces"
)

type TeacherUsecase struct {
	repo interfaces.TeacherRepository
}

func NewTeacherUsecase(repo interfaces.TeacherRepository) *TeacherUsecase {
	return &TeacherUsecase{repo: repo}
}

func (tu *TeacherUsecase) CreateTeacher(teacher *entities.Teacher) (string, error) {
	return tu.repo.CreateTeacher(teacher)
}

func (tu *TeacherUsecase) GetAllTeachers() ([]*entities.Teacher, error) {
	return tu.repo.GetAllTeachers()
}

func (tu *TeacherUsecase) GetTeacherByID(id string) (*entities.Teacher, error) {
	return tu.repo.GetTeacherByID(id)
}

func (tu *TeacherUsecase) UpdateTeacher(id string, teacher *entities.Teacher) error {
	return tu.repo.UpdateTeacher(id, teacher)
}

func (tu *TeacherUsecase) DeleteTeacher(id string) error {
	return tu.repo.DeleteTeacher(id)
}

