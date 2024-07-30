package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	sourceURL := "file://pkg/database/migrations"
	databaseURL := "postgres://root:123456@localhost:4444/education_db?sslmode=disable"
	
	m, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		log.Fatalf("migrate error: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("up error: %v", err)
	}

	log.Println("Migration completed")
}
