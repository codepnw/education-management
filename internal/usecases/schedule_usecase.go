package usecases

import (
	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/interfaces"
)

type ScheduleUsecase struct {
	repo interfaces.ScheduleRepository
}

func NewScheduleUsecase(repo interfaces.ScheduleRepository) *ScheduleUsecase {
	return &ScheduleUsecase{repo: repo}
}

func (su *ScheduleUsecase) CreateSchedule(req *entities.Schedule) (string, error) {
	return su.repo.CreateSchedule(req)
}

func (su *ScheduleUsecase) GetAllSchedule() ([]*entities.Schedule, error) {
	return su.repo.GetAllSchedule()
}

func (su *ScheduleUsecase) GetScheduleByID(id string) (*entities.Schedule, error) {
	return su.repo.GetScheduleByID(id)
}

func (su *ScheduleUsecase) UpdateSchedule(id string, req *entities.Schedule) error {
	return su.repo.UpdateSchedule(id, req)
}

func (su *ScheduleUsecase) DeleteSchedule(id string) error {
	return su.repo.DeleteSchedule(id)
}

