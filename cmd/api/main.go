package main

import (
	"log"
	"net/http"

	"github.com/conqdat/books-api/internal/config"
	"github.com/conqdat/books-api/internal/database"
	"github.com/conqdat/books-api/internal/handlers"
	repository "github.com/conqdat/books-api/internal/repository/postgres"
	"github.com/conqdat/books-api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to database
	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create table and sample data
	if err := database.CreateTable(db); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Initialize layers
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	// Setup router
	router := gin.Default()

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"message": "Books API is running",
		})
	})

	// Books API
	api := router.Group("/api/v1")
	{
		api.GET("/books", bookHandler.GetBooks)
	}

	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	log.Fatal(router.Run(":" + cfg.Server.Port))
}