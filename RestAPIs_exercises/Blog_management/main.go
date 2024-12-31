package main

import (
	"log"
	"net/http"

	"blog-management-system/handlers"
	"blog-management-system/middleware"
	"blog-management-system/database"
)

func main() {
	// Initialize the database
	db := database.InitDB()
	defer db.Close()

	// Set up routes
	http.HandleFunc("/blog", middleware.Logging(middleware.ValidateJSON(handlers.CreateBlog(db))))
	http.HandleFunc("/blog/", middleware.Logging(handlers.GetBlogByID(db)))
	http.HandleFunc("/blogs", middleware.Logging(handlers.GetAllBlogs(db)))
	http.HandleFunc("/blog/update/", middleware.Logging(middleware.ValidateJSON(handlers.UpdateBlog(db))))
	http.HandleFunc("/blog/delete/", middleware.Logging(handlers.DeleteBlog(db)))

	// Start server
	log.Println("Server running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
