package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	
	"strings"
	

	"blog-management-system/models"
)

func CreateBlog(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var blog models.Blog
		err := json.NewDecoder(r.Body).Decode(&blog)
		if err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}

		query := `INSERT INTO blogs (title, content, author) VALUES (?, ?, ?)`
		_, err = db.Exec(query, blog.Title, blog.Content, blog.Author)
		if err != nil {
			http.Error(w, "Failed to create blog", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Blog created successfully")
	}
}

func GetBlogByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/blog/")
		query := `SELECT * FROM blogs WHERE id = ?`
		row := db.QueryRow(query, id)

		var blog models.Blog
		err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
		if err == sql.ErrNoRows {
			http.Error(w, "Blog not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Failed to fetch blog", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(blog)
	}
}

func GetAllBlogs(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := `SELECT * FROM blogs`
		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, "Failed to fetch blogs", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var blogs []models.Blog
		for rows.Next() {
			var blog models.Blog
			err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
			if err != nil {
				http.Error(w, "Failed to fetch blogs", http.StatusInternalServerError)
				return
			}
			blogs = append(blogs, blog)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(blogs)
	}
}

func UpdateBlog(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		id := strings.TrimPrefix(r.URL.Path, "/blog/update/")
		var blog models.Blog
		err := json.NewDecoder(r.Body).Decode(&blog)
		if err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}

		query := `UPDATE blogs SET title = ?, content = ?, author = ? WHERE id = ?`
		_, err = db.Exec(query, blog.Title, blog.Content, blog.Author, id)
		if err != nil {
			http.Error(w, "Failed to update blog", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Blog updated successfully")
	}
}

func DeleteBlog(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		id := strings.TrimPrefix(r.URL.Path, "/blog/delete/")
		query := `DELETE FROM blogs WHERE id = ?`
		_, err := db.Exec(query, id)
		if err != nil {
			http.Error(w, "Failed to delete blog", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Blog deleted successfully")
	}
}
