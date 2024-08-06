package entities

import "time"

type Enrollment struct {
	ID           string    `json:"id"`
	StudentID    string    `json:"student_id"`
	CourseID     string    `json:"course_id"`
	DateEnrolled time.Time `json:"date_enrolled"`
	Grade        string    `json:"grade"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
