package services

import "time"

type Service struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"           binding:"required"`
	URL            string     `json:"url"            binding:"required,url"`
	Environment    string     `json:"environment"    binding:"required"`
	Status         string     `json:"status"`
	ResponseTimeMs int        `json:"response_time_ms"`
	LastStatusCode int        `json:"last_status_code"`
	LastCheckedAt  *time.Time `json:"last_checked_at"`
	CreatedAt      time.Time  `json:"created_at"`
}