package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/codepnw/education/internal/entities"
	"github.com/codepnw/education/internal/interfaces"
)

type scheduleRepository struct {
	db *sql.DB
}

func NewScheduleRepository(db *sql.DB) interfaces.ScheduleRepository {
	return &scheduleRepository{db: db}
}

func (sr *scheduleRepository) CreateSchedule(req *entities.Schedule) (string, error) {
	query := `
		INSERT INTO schedule (
			course_id,
			classroom_id,
			start_time,
			end_time
		) VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	err := sr.db.QueryRow(
		query,
		req.CourseID,
		req.ClassroomID,
		req.StartTime,
		req.EndTime,
	).Scan(&req.ID)

	if err != nil {
		return "", fmt.Errorf("query schedule failed: %v", err)
	}
	return req.ID, nil
}

func (sr *scheduleRepository) GetAllSchedule() ([]*entities.Schedule, error) {
	query := `
		SELECT
			id,
			course_id,
			classroom_id,
			start_time,
			end_time,
			created_at,
			updated_at
		FROM schedule;
	`
	rows, err := sr.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query schedule failed: %v", err)
	}
	defer rows.Close()

	schedule := make([]*entities.Schedule, 0)

	for rows.Next() {
		scd := new(entities.Schedule)

		if err := rows.Scan(
			scd.ID,
			scd.CourseID,
			scd.ClassroomID,
			scd.StartTime,
			scd.EndTime,
			scd.CreatedAt,
			scd.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan schedule failed: %v", err)
		}
		schedule = append(schedule, scd)
	}
	return schedule, nil
}

func (sr *scheduleRepository) GetScheduleByID(id string) (*entities.Schedule, error) {
	query := `
		SELECT	
			id,
			course_id,
			classroom_id,
			start_time,
			end_time,
			created_at,
			updated_at
		FROM schedule;
	`
	resp := new(entities.Schedule)

	if err := sr.db.QueryRow(query, id).Scan(
		resp.ID,
		resp.CourseID,
		resp.ClassroomID,
		resp.StartTime,
		resp.EndTime,
		resp.CreatedAt,
		resp.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("query schedule failed: %v", err)
	}
	return resp, nil
}

func (sr *scheduleRepository) UpdateSchedule(id string, req *entities.Schedule) error {
	query := `
		UPDATE schedule SET
			course_id = $1,
			classroom_id = $2,
			start_time = $3,
			end_time = $4,
		WHERE id = $5;
	`

	_, err := sr.db.Exec(
		query,
		req.CourseID,
		req.ClassroomID,
		req.StartTime,
		req.EndTime,
		id,
	)

	if err != nil {
		return fmt.Errorf("update schedule failed: %v", err)
	}
	return nil
}

func (sr *scheduleRepository) DeleteSchedule(id string) error {
	query := `DELETE FROM schedule WHERE id = $1;`

	_, err := sr.db.ExecContext(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("delete schedule failed: %v", err)
	}
	return nil
}
