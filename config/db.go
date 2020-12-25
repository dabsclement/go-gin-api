package config

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"

	controllers "go-gin-api/controllers"
)

// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "clem",
		Password: "postgres",
		Addr:     "localhost:5432",
		Database: "go-gin-api",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateTodoTable(db)
	controllers.InitiateDB(db)
	return db
}
