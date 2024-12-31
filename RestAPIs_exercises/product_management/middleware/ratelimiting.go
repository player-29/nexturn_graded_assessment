
package middleware



import (

    "github.com/gin-gonic/gin"

    "golang.org/x/time/rate"

)



func RateLimitingMiddleware(rps int) gin.HandlerFunc {

    limiter := rate.NewLimiter(rate.Limit(rps), rps)



    return func(c *gin.Context) {

        if !limiter.Allow() {

            c.AbortWithStatusJSON(429, gin.H{"error": "Too many requests"})

            return

        }

        c.Next()

    }

}
