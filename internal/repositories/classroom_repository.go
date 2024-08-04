package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/interfaces"
)

type classroomRepository struct {
	db *sql.DB
}

func NewClassroomRepository(db *sql.DB) interfaces.ClassroomRepository {
	return &classroomRepository{db: db}
}

func (cr *classroomRepository) CreateClassroom(req *entities.Classroom) (string, error) {
	query := `
		INSERT INTO classroom (
			name,
			location
		) VALUES ($1, $2)
		RETURNING id;
	`
	err := cr.db.QueryRow(query, req.Name, req.Location).Scan(&req.ID)
	if err != nil {
		return "", fmt.Errorf("query classroom failed: %v", err)
	}
	return req.ID, nil
}

func (cr *classroomRepository) GetAllClassroom() ([]*entities.Classroom, error) {
	query := `
		SELECT id, name, location, created_at, updated_at
		FROM classroom;
	`
	rows, err := cr.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query classroom failed: %v", err)
	}
	defer rows.Close()

	classroom := make([]*entities.Classroom, 0)

	for rows.Next() {
		var room entities.Classroom

		err := rows.Scan(
			&room.ID,
			&room.Name,
			&room.Location,
			&room.CreatedAt,
			&room.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan classroom failed: %v", err)
		}
		classroom = append(classroom, &room)
	}

	return classroom, nil
}

func (cr *classroomRepository) GetClassroomByID(id string) (*entities.Classroom, error) {
	query := `
		SELECT id, name, location, created_at, updated_at
		FROM classroom
		WHERE id = $1;
	`
	classroom := new(entities.Classroom)

	err := cr.db.QueryRow(query, id).Scan(
		&classroom.ID,
		&classroom.Name,
		&classroom.Location,
		&classroom.CreatedAt,
		&classroom.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("classroom not found: %v", err)
	}

	return classroom, nil
}

func (cr *classroomRepository) UpdateClassroom(id string, req *entities.Classroom) error {
	query := `
		UPDATE classroom SET
			name = $1,
			location = $2
		WHERE id = $3;
	`
	_, err := cr.db.Exec(query, req.Name, req.Location, id)
	if err != nil {
		return fmt.Errorf("update classroom failed: %v", err)
	}
	return nil
}

func (cr *classroomRepository) DeleteClassroom(id string) error {
	query := ` DELETE FROM classroom WHERE id = $1;`

	_, err := cr.db.ExecContext(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("delete classroom failed: %v", err)
	}
	return nil
}
