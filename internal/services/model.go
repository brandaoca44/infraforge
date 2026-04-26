package services

import "time"

type Service struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	URL            string     `json:"url"`
	Environment    string     `json:"environment"`
	Status         string     `json:"status"`
	ResponseTimeMs int        `json:"response_time_ms"`
	LastCheckedAt  *time.Time `json:"last_checked_at"`
	CreatedAt      time.Time  `json:"created_at"`
}