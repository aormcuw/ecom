package main

import (
	//+
	"log"
	"os"

	"github.com/aormcuw/ecom/config"
	"github.com/aormcuw/ecom/db"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	//+
)

func main() {
	gormDB, err := db.NewPostgresStorage(config.Envs.DBURL) //+
	if err != nil {
		log.Fatal("failed to connect to postgres ", err)
	}
	sqlDB, err := gormDB.DB() //+
	if err != nil {           //+
		log.Fatal("failed to get *sql.DB from *gorm.DB ", err) //+
	} //+
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{}) //+
	if err != nil {
		log.Fatal("failed to create migration instance ", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations/",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("failed to create migration engine ", err)
	}

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("failed to migrate up ", err)
		}
	} else if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("failed to migrate down ", err)
		}
	} else {
		log.Fatal("Invalid command. Use 'up' or 'down'")
	}
}
