package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/interfaces"
)

type studentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) interfaces.IStudentRepository {
	return &studentRepository{db: db}
}

func (sr *studentRepository) CreateStudent(student entities.Student) (string, error) {
	query := `
		INSERT INTO students (
			first_name,
			last_name,
			dob,
			phone,
			address
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	err := sr.db.QueryRow(
		query,
		student.FirstName,
		student.LastName,
		student.DOB,
		student.Phone,
		student.Address,
	).Scan(&student.ID)

	if err != nil {
		return "", fmt.Errorf("query student failed: %v", err)
	}

	return student.ID, nil
}

func (sr *studentRepository) GetAllStudents() ([]entities.Student, error) {
	query := `
		SELECT
			id,
			first_name,
			last_name,
			dob,
			phone,
			address,
			created_at,
			updated_at
		FROM students;
	`
	rows, err := sr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []entities.Student
	for rows.Next() {
		var student entities.Student

		err = rows.Scan(
			&student.ID,
			&student.FirstName,
			&student.LastName,
			&student.DOB,
			&student.Phone,
			&student.Address,
			&student.CreatedAt,
			&student.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}

func (sr *studentRepository) UpdateStudentByID(id string, student entities.Student) error {
	query := `
		UPDATE students SET 
			first_name = $1,
			last_name = $2,
			dob = $3,
			phone = $4,
			address = $5
		WHERE id = $6;
	`
	_, err := sr.db.Exec(
		query, 
		student.FirstName, 
		student.LastName, 
		student.DOB, 
		student.Phone, 
		student.Address,
		id,
	)

	if err != nil {
		return fmt.Errorf("update student failed: %v", err)
	}
	return nil
}

func (sr *studentRepository) DeleteStudent(id string) error {
	query := `DELETE FROM students WHERE id = $1;`

	_, err := sr.db.ExecContext(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("delete student failed: %v", err)
	}
	return nil
}
