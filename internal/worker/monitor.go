package worker

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/brandaoca44/infraforge/internal/database"
)

func StartMonitor() {
	go func() {
		for {
			checkServices()
			time.Sleep(10 * time.Second)
		}
	}()
}

func checkServices() {
	rows, err := database.DB.Query(context.Background(), `
		SELECT id, url FROM services
	`)
	if err != nil {
		log.Println("error fetching services:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var url string

		if err := rows.Scan(&id, &url); err != nil {
			continue
		}

		start := time.Now()
		resp, err := http.Get(url)
		duration := time.Since(start).Milliseconds()

		status := "offline"

		if err == nil && resp.StatusCode == 200 {
			status = "online"
		}

		_, err = database.DB.Exec(context.Background(), `
			UPDATE services
			SET status=$1, response_time_ms=$2, last_checked_at=NOW()
			WHERE id=$3
		`, status, duration, id)

		if err != nil {
			log.Println("error updating service:", err)
		}
	}
}