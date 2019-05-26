package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	tx "github.com/tax-calculator/internal/data/tax"
	"github.com/tax-calculator/internal/httpserver"
	"github.com/tax-calculator/internal/tax"

	"github.com/tax-calculator/config"
)

func main() {
	config, err := config.GetConfiguration()
	if err != nil {
		log.Fatalln("failed to get configuration: ", err)
	}

	db, err := sql.Open("postgres", config.DBConnectionString)
	if err != nil {
		log.Fatalln("failed when open DB: ", err)
	}

	service := tx.NewService(db)
	taxService := tax.NewService(service)

	s := httpserver.NewServer(taxService)

	s.Serve()
}
