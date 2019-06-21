package routing

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetPingRoute - simple string with current time.
func GetPingRoute(ctx *gin.Context) {
	response := fmt.Sprintf("Works! Now is: %s", time.Now())
	ctx.String(http.StatusOK, response)
}

// PostPingRoute - simple JSON with current time.
func PostPingRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Works! Now is: %s", time.Now()),
	})
}
