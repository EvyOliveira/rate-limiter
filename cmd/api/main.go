package api

import (
	"net/http"
	"time"

	"github.com/EvyOliveira/rate-limiter/pkg/ratelimit"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	rateLimit := ratelimit.NewRateLimit(100, 5*time.Second)

	app.Use(rateLimit.Apply())

	app.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})
	app.Run()
}
