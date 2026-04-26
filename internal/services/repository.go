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