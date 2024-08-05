package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/interfaces"
)

type courseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) interfaces.CourseRepository {
	return &courseRepository{db: db}
}

func (cr *courseRepository) CreateCourse(req *entities.Course) (string, error) {
	query := `
		INSERT INTO courses (
			title,
			description,
			teacher_id
		) VALUES ($1, $2, $3)
		RETURNING id;
	`
	err := cr.db.QueryRow(
		query,
		req.Title,
		req.Description,
		req.TeacherID,
	).Scan(&req.ID)

	if err != nil {
		return "", fmt.Errorf("query course failed: %v", err)
	}
	return req.ID, nil
}

func (cr *courseRepository) GetAllCourses() ([]*entities.Course, error) {
	query := `
		SELECT
			id, 
			title, 
			description,
			teacher_id,
			created_at,
			updated_at
		FROM courses;
	`
	courses := make([]*entities.Course, 0)

	rows, err := cr.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query courses failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var course entities.Course
		err := rows.Scan(
			&course.ID,
			&course.Title,
			&course.Description,
			&course.TeacherID,
			&course.CreatedAt,
			&course.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("scan courses failed: %v", err)
		}
		courses = append(courses, &course)
	}
	return courses, nil
}

func (cr *courseRepository) GetCourseByID(id string) (*entities.Course, error) {
	query := `
		SELECT
			id, 
			title, 
			description,
			teacher_id,
			created_at,
			updated_at
		FROM courses
		WHERE id = $1;
	`
	course := new(entities.Course)

	err := cr.db.QueryRow(query, id).Scan(
		&course.ID,
		&course.Title,
		&course.Description,
		&course.TeacherID,
		&course.CreatedAt,
		&course.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("query course failed: %v", err)
	}
	return course, nil
}

func (cr *courseRepository) UpdateCourse(id string, req *entities.Course) error {
	query := `
		UPDATE courses SET
			title = $1
			description = $2
			teacher_id = $3
		WHERE id = $4;
	`
	_, err := cr.db.Exec(
		query,
		req.Title,
		req.Description,
		req.TeacherID,
		id,
	)

	if err != nil {
		return fmt.Errorf("query course failed: %v", err)
	}
	return nil
}

func (cr *courseRepository) DeleteCourse(id string) error {
	query := `DELETE FROM courses WHERE id = $1;`

	_, err := cr.db.ExecContext(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("query course failed: %v", err)
	}
	return nil
}
