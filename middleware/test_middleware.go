package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Printnumber 印出123
func Printnumber(c *gin.Context) {
	start := time.Now()
	c.Set("startTime", start)
	// Pass on to the next-in-chain

	c.Next()
}
