// package main

// import (
// 	db "go-sqlite-crud-product/config"
// 	"go-sqlite-crud-product/controller"
// 	"go-sqlite-crud-product/repository"
// 	"go-sqlite-crud-product/service"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// Initialize database connection
// 	db.InitializeDatabase()

// 	// Create repository, service, and controller
// 	userRepo := repository.NewUserRepository(db.GetDB())
// 	userService := service.NewUserService(userRepo)
// 	userController := controller.NewUserController(userService)

// 	// Create repository, service, and controller for Products
// 	productRepo := repository.NewProductRepository(db.GetDB())
// 	productService := service.NewProductService(productRepo)
// 	productController := controller.NewProductController(productService)

// 	// Initialize Gin router
// 	r := gin.Default()

// 	// Routes
// 	r.POST("/users", userController.CreateUser)
// 	r.GET("/users/:id", userController.GetUser)
// 	r.GET("/users", userController.GetAllUsers)
// 	r.PUT("/users/:id", userController.UpdateUser)
// 	r.DELETE("/users/:id", userController.DeleteUser)

// 	// Product Routes
// 	r.POST("/products", productController.CreateProduct)
// 	r.GET("/products/:id", productController.GetProduct)
// 	r.GET("/products", productController.GetAllProducts)
// 	r.PUT("/products/:id", productController.UpdateProduct)
// 	r.DELETE("/products/:id", productController.DeleteProduct)

// 	// Start server
// 	r.Run(":8081")

// }


package main

import (
	db "go-sqlite-crud-product/config"
	"go-sqlite-crud-product/controller"
	"go-sqlite-crud-product/middleware"
	"go-sqlite-crud-product/repository"
	"go-sqlite-crud-product/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db.InitializeDatabase()

	// Create repository, service, and controller
	userRepo := repository.NewUserRepository(db.GetDB())
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	productRepo := repository.NewProductRepository(db.GetDB())
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	// Initialize Gin router
	r := gin.Default()

	// Global Middleware
	r.Use(middleware.LoggingMiddleware())

	// Routes
	userRoutes := r.Group("/users")
	{
		userRoutes.Use(middleware.JWTAuthMiddleware())
		userRoutes.POST("", middleware.InputValidationMiddleware([]string{"name", "email"}), userController.CreateUser)
		userRoutes.GET("/:id", userController.GetUser)
		userRoutes.GET("", userController.GetAllUsers)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}

	productRoutes := r.Group("/products")
	{
		productRoutes.Use(middleware.RateLimitingMiddleware(5)) // 5 requests per minute
		productRoutes.POST("", middleware.InputValidationMiddleware([]string{"name", "price"}), productController.CreateProduct)
		productRoutes.GET("/:id", productController.GetProduct)
		productRoutes.GET("", productController.GetAllProducts)
		productRoutes.PUT("/:id", productController.UpdateProduct)
		productRoutes.DELETE("/:id", productController.DeleteProduct)
	}

	// Start server
	r.Run(":8080")
}
