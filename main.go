package main

import (
	"FinalProject4/config"
	"FinalProject4/controller"
	"FinalProject4/middleware"
	"FinalProject4/model/entity"
	"FinalProject4/model/seed"
	"FinalProject4/repository"
	"FinalProject4/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db := config.StartDB()
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Category{})
	db.AutoMigrate(&entity.Product{})
	db.AutoMigrate(&entity.TransactionHistory{})

	db.Create(&seed.User)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService, productService)

	transactionHistoryRepository := repository.NewTransactionHistoryRepository(db)
	transactionHistoryService := service.NewTransactionHistoryService(transactionHistoryRepository, productRepository, userRepository)
	transactionHistoryController := controller.NewTransactionHistoryController(transactionHistoryService, productService)

	router := gin.Default()

	// Users
	router.POST("/users/register", userController.RegisterUser)
	router.POST("/users/login", userController.LoginUser)
	router.PUT("/users/topup", middleware.AuthMiddleware(), userController.UserTopup)

	// Category
	router.POST("/categories", middleware.AuthMiddleware(), middleware.IsAdmin(), categoryController.CreateCategory)
	router.GET("/categories", middleware.AuthMiddleware(), middleware.IsAdmin(), categoryController.GetCategory)
	router.PATCH("/categories/:id", middleware.AuthMiddleware(), middleware.IsAdmin(), categoryController.UpdateCategory)
	router.DELETE("/categories/:id", middleware.AuthMiddleware(), middleware.IsAdmin(), categoryController.DeleteCategory)

	// Product
	router.POST("/products", middleware.AuthMiddleware(), middleware.IsAdmin(), productController.CreateProduct)
	router.GET("/products", productController.GetProduct)
	router.PUT("/products/:id", middleware.AuthMiddleware(), middleware.IsAdmin(), productController.UpdateProduct)
	router.DELETE("/products/:id", middleware.AuthMiddleware(), middleware.IsAdmin(), productController.DeleteProduct)

	// Transaction History
	router.POST("/transactionhistories", middleware.AuthMiddleware(), transactionHistoryController.CreateTransactionHistory)
	router.GET("/transactionhistories/my-transactions", middleware.AuthMiddleware(), transactionHistoryController.GetMyTransactionHistory)
	router.GET("/transactionhistories/user-transactions", middleware.AuthMiddleware(), middleware.IsAdmin(), transactionHistoryController.GetAllTransactionHistory)

	router.Run()
}
