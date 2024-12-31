package middleware

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// AuthMiddleware validates the user credentials with basic auth
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("AuthMiddleware: Started %s %s\n", r.Method, r.URL.Path)

		// Get Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			http.Error(w, "Unauthorized: Missing or invalid Authorization header", http.StatusUnauthorized)
			fmt.Printf("AuthMiddleware: Missing or invalid Authorization header. Completed %s in %v\n", r.URL.Path, time.Since(start))
			return
		}

		// Decode credentials from the header
		payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "Basic "))
		if err != nil {
			http.Error(w, "Unauthorized: Invalid Base64 encoding", http.StatusUnauthorized)
			fmt.Printf("AuthMiddleware: Invalid Base64 encoding. Completed %s in %v\n", r.URL.Path, time.Since(start))
			return
		}
		credentials := strings.SplitN(string(payload), ":", 2)
		if len(credentials) != 2 {
			http.Error(w, "Unauthorized: Invalid credentials format", http.StatusUnauthorized)
			fmt.Printf("AuthMiddleware: Invalid credentials format. Completed %s in %v\n", r.URL.Path, time.Since(start))
			return
		}

		// Extract username and password
		username := credentials[0]
		password := credentials[1]

		// Connect to SQLite database
		db, err := sql.Open("sqlite3", "users.db")
		if err != nil {
			http.Error(w, "Internal Server Error: Database connection failed", http.StatusInternalServerError)
			fmt.Printf("AuthMiddleware: Database connection failed. Completed %s in %v\n", r.URL.Path, time.Since(start))
			return
		}
		defer db.Close()

		// Query the database for the user
		var dbPassword string
		query := "SELECT password FROM users WHERE username = ?"
		err = db.QueryRow(query, username).Scan(&dbPassword)
		if err != nil || dbPassword != password {
			http.Error(w, "Unauthorized: Invalid credentials", http.StatusUnauthorized)
			fmt.Printf("AuthMiddleware: Invalid credentials. Completed %s in %v\n", r.URL.Path, time.Since(start))
			return
		}

		fmt.Printf("AuthMiddleware: Successfully authenticated user '%s'. Completed %s in %v\n", username, r.URL.Path, time.Since(start))
		next.ServeHTTP(w, r)
	})
}



// LoggingMiddleware logs the details of each request

func LoggingMiddleware() gin.HandlerFunc {

    return func(c *gin.Context) {

        startTime := time.Now()



        // Process request

        c.Next()



        // Log request details

        endTime := time.Now()

        latency := endTime.Sub(startTime)

        statusCode := c.Writer.Status()

        clientIP := c.ClientIP()

        method := c.Request.Method

        path := c.Request.URL.Path



        log.Printf("%s %s %s %d %s", clientIP, method, path, statusCode, latency)

    }

}

