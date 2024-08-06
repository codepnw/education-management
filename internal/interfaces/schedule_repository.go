package interfaces

import "github.com/codepnw/education/internal/entities"

type ScheduleRepository interface {
	CreateSchedule(req *entities.Schedule) (string, error)
	GetAllSchedule() ([]*entities.Schedule, error)
	GetScheduleByID(id string) (*entities.Schedule, error)
	UpdateSchedule(id string, req *entities.Schedule) error
	DeleteSchedule(id string) error
}