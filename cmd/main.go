package main

import (
	"log"

	"github.com/aormcuw/ecom/cmd/api"
	"github.com/aormcuw/ecom/db"
)

func main() {
	db, err := db.NewPostgresStorage()
	if err != nil {
		log.Fatal("failed to connect to postgres ", err)
	}
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal("Error running")
	}

}
