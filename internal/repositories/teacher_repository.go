package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/interfaces"
)

type teacherRepository struct {
	db *sql.DB
}

func NewTeacherRepository(db *sql.DB) interfaces.TeacherRepository {
	return &teacherRepository{db: db}
}

func (tr *teacherRepository) CreateTeacher(teacher *entities.Teacher) (string, error) {
	query := `
		INSERT INTO teachers (
			first_name,
			last_name,
			department,
			dob,
			phone,
			address
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`
	err := tr.db.QueryRow(
		query,
		teacher.FirstName,
		teacher.LastName,
		teacher.Department,
		teacher.DOB,
		teacher.Phone,
		teacher.Address,
	).Scan(&teacher.ID)

	if err != nil {
		return "", fmt.Errorf("query teacher failed: %v", err)
	}

	return teacher.ID, nil
}

func (tr *teacherRepository) GetAllTeachers() ([]*entities.Teacher, error) {
	query := `
		SELECT
			id,
			first_name,
			last_name,
			department,
			dob,
			phone,
			address,
			created_at,
			updated_at
		FROM teachers;
	`
	rows, err := tr.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query teachers failed: %v", err)
	}
	defer rows.Close()

	teachers := make([]*entities.Teacher, 0)

	for rows.Next() {
		var teacher entities.Teacher

		err = rows.Scan(
			&teacher.ID,
			&teacher.FirstName,
			&teacher.LastName,
			&teacher.Department,
			&teacher.DOB,
			&teacher.Phone,
			&teacher.Address,
			&teacher.CreatedAt,
			&teacher.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan teachers failed: %v", err)
		}
		teachers = append(teachers, &teacher)
	}

	return teachers, nil
}

func (tr *teacherRepository) GetTeacherByID(id string) (*entities.Teacher, error) {
	query := `
		SELECT
			id,
			first_name,
			last_name,
			department,
			dob,
			phone,
			address,
			created_at,
			updated_at
		FROM teachers
		WHERE id = $1;		
	`
	teacher := new(entities.Teacher)

	err := tr.db.QueryRow(query, id).Scan(
		&teacher.ID,
		&teacher.FirstName,
		&teacher.LastName,
		&teacher.Department,
		&teacher.DOB,
		&teacher.Phone,
		&teacher.Address,
		&teacher.CreatedAt,
		&teacher.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("teacher not found: %v", err)
	}

	return teacher, nil
}

func (tr *teacherRepository) UpdateTeacher(id string, teacher *entities.Teacher) error {
	query := `
		UPDATE teachers SET 
			first_name = $1,
			last_name = $2,
			department = $3,
			dob = $4,
			phone = $5,
			address = $6
		WHERE id = $7;
	`
	_, err := tr.db.Exec(
		query,
		teacher.FirstName,
		teacher.LastName,
		teacher.Department,
		teacher.DOB,
		teacher.Phone,
		teacher.Address,
		id,
	)

	if err != nil {
		return fmt.Errorf("update teacher failed: %v", err)
	}
	return nil
}

func (tr *teacherRepository) DeleteTeacher(id string) error {
	query := `DELETE FROM teachers WHERE id = $1;`

	_, err := tr.db.ExecContext(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("delete teacher failed: %v", err)
	}
	return nil
}
