package worker

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/brandaoca44/infraforge/internal/config"
	"github.com/brandaoca44/infraforge/internal/database"
)

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

func StartMonitor(ctx context.Context, cfg config.Config) {
	go func() {
		ticker := time.NewTicker(cfg.MonitorInterval)
		defer ticker.Stop()

		log.Println("Service monitor started")

		for {
			select {
			case <-ticker.C:
				checkServices(ctx)
			case <-ctx.Done():
				log.Println("Service monitor stopped")
				return
			}
		}
	}()
}

func checkServices(ctx context.Context) {
	rows, err := database.DB.Query(ctx, `
		SELECT id, url FROM services
	`)
	if err != nil {
		log.Println("[ERROR] fetching services:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var url string

		if err := rows.Scan(&id, &url); err != nil {
			log.Println("[WARN] scanning service row:", err)
			continue
		}

		start := time.Now()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			log.Printf("[WARN] creating request service_id=%s url=%s: %v", id, url, err)
			continue
		}

		resp, err := httpClient.Do(req)
		duration := time.Since(start).Milliseconds()

		status := "offline"
		statusCode := 0

		if err == nil {
			statusCode = resp.StatusCode
			resp.Body.Close()

			if resp.StatusCode >= 200 && resp.StatusCode < 400 {
				status = "online"
			}
		} else {
			log.Printf("[WARN] pinging service_id=%s url=%s: %v", id, url, err)
		}

		_, err = database.DB.Exec(ctx, `
			UPDATE services
			SET status=$1, response_time_ms=$2, last_checked_at=NOW(), last_status_code=$3
			WHERE id=$4
		`, status, duration, statusCode, id)

		if err != nil {
			log.Printf("[ERROR] updating service_id=%s: %v", id, err)
		}
	}

	if err := rows.Err(); err != nil {
		log.Println("[ERROR] reading service rows:", err)
	}
}