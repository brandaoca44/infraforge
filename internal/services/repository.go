package services

import (
	"context"

	"github.com/brandaoca44/infraforge/internal/database"
)

func CreateService(s Service) error {
	query := `
		INSERT INTO services (name, url, environment)
		VALUES ($1, $2, $3)
	`

	_, err := database.DB.Exec(context.Background(), query, s.Name, s.URL, s.Environment)
	return err
}

func GetServices() ([]Service, error) {
	rows, err := database.DB.Query(context.Background(), `
		SELECT id, name, url, environment, status, response_time_ms, last_status_code, last_checked_at, created_at
		FROM services
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	services := []Service{}

	for rows.Next() {
		var s Service

		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.URL,
			&s.Environment,
			&s.Status,
			&s.ResponseTimeMs,
			&s.LastStatusCode,
			&s.LastCheckedAt,
			&s.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		services = append(services, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return services, nil
}

func DeleteService(id string) (bool, error) {
	result, err := database.DB.Exec(context.Background(), `
		DELETE FROM services WHERE id=$1
	`, id)
	if err != nil {
		return false, err
	}

	return result.RowsAffected() > 0, nil
}