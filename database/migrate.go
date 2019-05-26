package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"

	"github.com/tax-calculator/config"
)

func main() {
	config, err := config.GetConfiguration()
	if err != nil {
		log.Fatal("error when getting configuration: ", err)
	}

	db, err := sql.Open("postgres", config.DBConnectionString)
	if err != nil {
		log.Fatal("error when open postgres connection: ", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("error when creating postgres instance: ", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres", driver)

	if err != nil {
		log.Fatal("error when creating database instance: ", err)
	}

	if err := m.Up(); err != nil {
		log.Fatal("error when migrate up: ", err)
	}

}
