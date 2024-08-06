package usecases

import (
	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/interfaces"
)

type EnrollmentUsecase struct {
	repo interfaces.EnrollmentRepository
}

func NewEnrollmentUsecase(repo interfaces.EnrollmentRepository) *EnrollmentUsecase {
	return &EnrollmentUsecase{repo: repo}
}

func (eu *EnrollmentUsecase) CreateEnroll(req *entities.Enrollment) (string, error) {
	return eu.repo.CreateEnroll(req)
}

func (eu *EnrollmentUsecase) GetAllEnroll() ([]*entities.Enrollment, error) {
	return eu.repo.GetAllEnroll()
}

func (eu *EnrollmentUsecase) GetEnrollByID(id string) (*entities.Enrollment, error) {
	return eu.repo.GetEnrollByID(id)
}

func (eu *EnrollmentUsecase) UpdateEnroll(id string, req *entities.Enrollment) error {
	return eu.repo.UpdateEnroll(id, req)
}

func (eu *EnrollmentUsecase) DeleteEnroll(id string) error {
	return eu.repo.DeleteEnroll(id)
}

