package main

import (
	"database/sql"
	"fmt"
	"os"
	"practice/controllers"
	"practice/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	// Get database configuration from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Construct database connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Connect to the database
	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("DB Connection Failed:", err)
		os.Exit(1)
	}
	defer DB.Close()

	// Check database connection
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Ping Failed:", err)
		os.Exit(1)
	}

	fmt.Println("DB Connection Success")

	// Run database migrations
	database.DbMigrate(DB)

	// Set up the router with Gin
	router := gin.Default()
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/persons", controllers.InsertPerson)
	router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/persons/:id", controllers.DeletePerson)

	// Start the server
	router.Run(":8080")
}
