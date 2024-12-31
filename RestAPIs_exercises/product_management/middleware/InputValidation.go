
package middleware



import (

    "github.com/gin-gonic/gin"

    "net/http"

)



// InputValidationMiddleware validates required fields in the request body

func InputValidationMiddleware(requiredFields []string) gin.HandlerFunc {

    return func(c *gin.Context) {

        var json map[string]interface{}

        if err := c.ShouldBindJSON(&json); err != nil {

            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})

            c.Abort()

            return

        }



        for _, field := range requiredFields {

            if _, ok := json[field]; !ok {

                c.JSON(http.StatusBadRequest, gin.H{"error": field + " is required"})

                c.Abort()

                return

            }

        }



        c.Next()

    }

}
