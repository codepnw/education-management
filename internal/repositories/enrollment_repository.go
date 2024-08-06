package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/interfaces"
)

type enrollmentRepository struct {
	db *sql.DB
}

func NewEnrollmentRepository(db *sql.DB) interfaces.Enrollment {
	return &enrollmentRepository{db: db}
}

func (er *enrollmentRepository) CreateEnroll(req *entities.Enrollment) (string, error) {
	query := `
		INSERT INTO enrollment (
			student_id,
			course_id,
			date_enrolled,
			grade
		) VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	err := er.db.QueryRow(
		query, 
		req.StudentID, 
		req.CourseID, 
		req.DateEnrolled, 
		req.Grade,
	).Scan(&req.ID)

	if err != nil {
		return "", fmt.Errorf("query enrollment failed: %v", err)
	}
	return req.ID, nil
}

func (er *enrollmentRepository) GetAllEnroll() ([]*entities.Enrollment, error) {
	query := `
		SELECT 
			id,
			student_id,
			course_id,
			date_enrolled,
			grade,
			created_at,
			updated_at
		FROM enrollment;
	`
	rows, err := er.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query enrollment failed: %v", err)
	}
	defer rows.Close()

	enrollments := make([]*entities.Enrollment, 0)

	for rows.Next() {
		enrollment := new(entities.Enrollment)
		err := rows.Scan(
			&enrollment.ID,
			&enrollment.StudentID,
			&enrollment.CourseID,
			&enrollment.DateEnrolled,
			&enrollment.Grade,
			&enrollment.CreatedAt,
			&enrollment.UpdatedAt,
		)
		
		if err != nil {
			return nil, fmt.Errorf("scan enrollment failed: %v", err)
		}
		enrollments = append(enrollments, enrollment)	
	}
	return enrollments, nil
}

func (er *enrollmentRepository) GetEnrollByID(id string) (*entities.Enrollment, error) {
	query := `
		SELECT 
			id,
			student_id,
			course_id,
			date_enrolled,
			grade,
			created_at,
			updated_at			
		FROM enrollment;
		WHERE id = $1;
	`
	enrollment := new(entities.Enrollment)

	err := er.db.QueryRow(query, id).Scan(
		&enrollment.ID,
		&enrollment.StudentID,
		&enrollment.CourseID,
		&enrollment.DateEnrolled,
		&enrollment.Grade,
		&enrollment.CreatedAt,
		&enrollment.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("query enrollment failed: %v", err)
	}
	return enrollment, nil
}

func (er *enrollmentRepository) UpdateEnroll(id string, req *entities.Enrollment) error {
	query := `
		UPDATE enrollment SET
			student_id,
			course_id,
			date_enrolled,
			grade
		WHERE id = $1;
	`
	_, err := er.db.Exec(query, req.StudentID, req.CourseID, req.DateEnrolled, req.Grade, id)
	if err != nil {
		return fmt.Errorf("update enrollment failed: %v", err)
	}
	return nil
}

func (er *enrollmentRepository) DeleteEnroll(id string) error {
	query := `DELETE FROM enrollment WHERE id = $1;`
	
	_, err := er.db.ExecContext(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("delete enrollment failed: %v", err)
	}
	return nil
}
