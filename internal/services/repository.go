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
		SELECT id, name, url, environment, status, response_time_ms, last_checked_at, created_at
		FROM services
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []Service

	for rows.Next() {
		var s Service

		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.URL,
			&s.Environment,
			&s.Status,
			&s.ResponseTimeMs,
			&s.LastCheckedAt,
			&s.CreatedAt,
		)
		if err != nil {
			continue
		}

		services = append(services, s)
	}

	return services, nil
}