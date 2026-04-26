package health

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"service":   "infraforge-api",
		"timestamp": time.Now().UTC(),
	})
}