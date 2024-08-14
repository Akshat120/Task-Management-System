package db

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
)

func InitDB() *pg.DB {
	// Get the connection string from the environment variable
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Parse the connection string and set up options for pg.Connect
	opts, err := pg.ParseURL(connStr)
	if err != nil {
		log.Fatalf("Error parsing DATABASE_URL: %v", err)
	}

	// Connect to the database
	db := pg.Connect(opts)

	// Check if the connection is successful by running a simple query
	if _, err := db.Exec("SELECT 1"); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Connected to the database successfully")
	return db
}
