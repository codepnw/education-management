package usecases

import (
	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/interfaces"
)

type CourseUsecase struct {
	repo interfaces.CourseRepository
}

func NewCourseUsecase(repo interfaces.CourseRepository) *CourseUsecase {
	return &CourseUsecase{repo: repo}
}

func (cu *CourseUsecase) CreateCourse(req *entities.Course) (string, error) {
	return cu.repo.CreateCourse(req)
}

func (cu *CourseUsecase) GetAllCourses() ([]*entities.Course, error) {
	return cu.repo.GetAllCourses()
}

func (cu *CourseUsecase) GetCourseByID(id string) (*entities.Course, error) {
	return cu.repo.GetCourseByID(id)
}

func (cu *CourseUsecase) UpdateCourse(id string, req *entities.Course) error {
	return cu.repo.UpdateCourse(id, req)
}

func (cu *CourseUsecase) DeleteCourse(id string) error {
	return cu.repo.DeleteCourse(id)
}

