package main

import (
	"log"

	"github.com/aormcuw/ecom/cmd/api" // Import the missing config package//+
	"github.com/aormcuw/ecom/config"
	"github.com/aormcuw/ecom/db"
	"gorm.io/gorm"
)

func main() {

	db, err := db.NewPostgresStorage(config.Envs.DBURL)
	if err != nil {
		log.Fatal("failed to connect to postgres ", err)
	}
	initStorage(db)
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal("Error running")
	}
	// -
}

func initStorage(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to connect to database ", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("failed to ping database ", err)
	}

	log.Println("Database connection successfully")

}
