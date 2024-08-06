package entities

import "time"

type Schedule struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"course_id"`
	ClassroomID string    `json:"classroom_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
