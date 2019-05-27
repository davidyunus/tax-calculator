package tax_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/davidyunus/tax-calculator/internal/tax"

	_ "github.com/lib/pq"

	"github.com/davidyunus/tax-calculator/config"
	tx "github.com/davidyunus/tax-calculator/internal/data/tax"
)

func TestCreate(t *testing.T) {
	config, err := config.GetConfiguration()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when get configuration", err)
	}

	connection := config.DBConnectionString
	db, err := sql.Open("postgres", connection)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening DB connection", err)
	}

	taxService := initService(db)

	defer db.Close()
	defer db.Exec(`DELETE FROM "tax" where "name" = 'pizza'`)

	err = taxService.Create(context.Background(), "pizza", 1, 2000)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when create tax data", err)
	}
}

func TestFindAll(t *testing.T) {
	config, err := config.GetConfiguration()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when get configuration", err)
	}

	connection := config.DBConnectionString
	db, err := sql.Open("postgres", connection)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening DB connection", err)
	}

	taxService := initService(db)

	defer db.Close()
	db.Exec(`INSERT INTO "tax"
		("name", "taxCode", "price", "createdAt", "updatedAt")
			VALUES 
		('pizza', 1, 2000, 2019-05-27 09:57:25, 2019-05-27 09:57:25)`)
	defer db.Exec(`DELETE FROM "tax" where "name" = 'pizza'`)

	taxData, err := taxService.FindAll(context.Background())
	if err != nil {
		t.Fatalf("an error '%s' was not expected when find all tax data", err)
	}

	assertEqual(t, len(taxData) > 0, true)
}

func TestFindByKeys(t *testing.T) {
	config, err := config.GetConfiguration()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when get configuration", err)
	}

	connection := config.DBConnectionString
	db, err := sql.Open("postgres", connection)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening DB connection", err)
	}

	taxService := initService(db)

	defer db.Close()
	db.Exec(`INSERT INTO "tax"
		("name", "taxCode", "price", "createdAt", "updatedAt")
			VALUES 
		('pizza', 1, 2000, 2019-05-27 09:57:25, 2019-05-27 09:57:25)`)
	defer db.Exec(`DELETE FROM "tax" where "name" = 'pizza'`)

	taxData, err := taxService.FindByKeys(context.Background(), 1)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when find all tax data", err)
	}
	assertEqual(t, taxData.TaxID > 0, true)
}

func TestFindByQuery(t *testing.T) {
	config, err := config.GetConfiguration()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when get configuration", err)
	}

	connection := config.DBConnectionString
	db, err := sql.Open("postgres", connection)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening DB connection", err)
	}

	taxService := initService(db)

	defer db.Close()
	db.Exec(`INSERT INTO "tax"
		("name", "taxCode", "price", "createdAt", "updatedAt")
			VALUES 
		('pizza', 1, 2000, 2019-05-27 09:57:25, 2019-05-27 09:57:25)`)
	defer db.Exec(`DELETE FROM "tax" where "name" = 'pizza'`)

	_, err = taxService.FindByQuery(context.Background(), "pizza")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when find all tax data", err)
	}

}

func TestUpdate(t *testing.T) {
	config, err := config.GetConfiguration()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when get configuration", err)
	}

	connection := config.DBConnectionString
	db, err := sql.Open("postgres", connection)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening DB connection", err)
	}

	taxService := initService(db)

	defer db.Close()
	db.Exec(`INSERT INTO "tax"
		("name", "taxCode", "price", "createdAt", "updatedAt")
			VALUES 
		('pizza', 1, 2000, 2019-05-27 09:57:25, 2019-05-27 09:57:25)`)
	defer db.Exec(`DELETE FROM "tax" where "name" = 'pizza'`)

	err = taxService.Update(context.Background(), 1, "pizza", 1, 2000)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when find all tax data", err)
	}

}

func assertEqual(t *testing.T, actual, expected interface{}) {
	if expected != actual {
		t.Fatalf("invalid value, expected: %v, got: %v", expected, actual)
	}
}

func initService(db *sql.DB) *tax.Service {
	service := tx.NewService(db)
	taxService := tax.NewService(service)
	return taxService
}
