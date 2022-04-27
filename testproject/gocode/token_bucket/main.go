package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

func RateLimitMiddleware(fillInterval time.Duration, cap int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			fmt.Println("-----------被限流了--------------",time.Now().Unix())
			c.AbortWithStatus(504)
			return
		}
		fmt.Println("通过限流",time.Now().Unix())
		c.Next()
	}

}

func main() {
	r := gin.New()
	fillInterval := 200 * time.Second
	r.GET("/", RateLimitMiddleware(fillInterval, 2), func(c *gin.Context) {
		c.String(http.StatusOK, "this is main")
	})
	r.Run(":8080")
}