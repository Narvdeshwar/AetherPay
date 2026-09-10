package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RateLimiterMiddleware(rdb *redis.Client, limit int64, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		// 1. Identity decide karo: Agar context me tenant_id hai toh wo lo, warna Client IP
		var identifier string
		if tenantID, exits := c.Get("tenant_id"); exits {
			identifier = fmt.Sprintf("rate:tenant:%v", tenantID)
		} else {
			identifier = fmt.Sprintf("rate:clientIP:%v", c.ClientIP())
		}
		// 2. Redis INCR (Atomically count badhao)
		count, err := rdb.Incr(ctx, identifier).Result()
		if err != nil {
			// Redis fail hone par traffic block na ho (Fail Open strategy)
			c.Next()
			return
		}
		// 3. Agar pehli request hai (count == 1), toh Expiration TTL set karo
		if count == 1 {
			rdb.Expire(ctx, identifier, window)
		}
		// 4. Check if limit exceeded
		if count > limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":       "Too many requests",
				"message":     "Rate limit exceeded. Please slow down.",
				"retry_after": window.String(),
			})
			return
		}
		c.Next()
	}
}
