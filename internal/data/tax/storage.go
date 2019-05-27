package tax

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/davidyunus/tax-calculator/internal/data"
)

// Storage ...
type Storage struct {
	queryable data.Queryable
}

// NewStorage ...
func NewStorage(q data.Queryable) *Storage {
	return &Storage{
		queryable: q,
	}
}

type rowScanner interface {
	Scan(dest ...interface{}) error
}

func selectQuery() string {
	return `SELECT "taxId", "name", "taxCode", "price", "createdAt", "updatedAt", "deletedAt"
		FROM "tax"`
}

func insertQuery(tax *Tax) (string, []interface{}) {
	return `INSERT INTO "tax"
		("name", "taxCode", "price", "createdAt", "updatedAt")
		VALUES 
		($1, $2, $3, $4, $5)
		RETURNING
		"taxId", "name", "taxCode", "price", "createdAt", "updatedAt", "deletedAt"`,
		[]interface{}{
			tax.Name, tax.TaxCode, tax.Price, tax.CreatedAt, tax.UpdatedAt,
		}
}

func updateQuery(tax *Tax) (string, []interface{}) {
	return `UPDATE "tax"
		SET 
			"name" = $1, "taxCode" = $2, "price" = $3, "updatedAt" = $4
		WHERE "taxId" = $5
		RETURNING 
		"taxId", "name", "taxCode", "price", "createdAt", "updatedAt", "deletedAt"`,
		[]interface{}{
			tax.Name, tax.TaxCode, tax.Price, tax.UpdatedAt, tax.TaxID,
		}
}

func scanTax(row rowScanner, tax *Tax) error {
	return row.Scan(&tax.TaxID, &tax.Name, &tax.TaxCode,
		&tax.Price, &tax.CreatedAt, &tax.UpdatedAt, &tax.DeletedAt)
}

// Create ...
func (s *Storage) Create(ctx context.Context, tax *Tax) error {
	queryable, ok := data.QueryableFromContext(ctx)
	if !ok {
		queryable = s.queryable
	}

	query, params := insertQuery(tax)
	row := queryable.QueryRow(query, params...)

	err := scanTax(row, tax)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

// FindAll ...
func (s *Storage) FindAll(ctx context.Context) ([]*Tax, error) {
	queryable, ok := data.QueryableFromContext(ctx)
	if !ok {
		queryable = s.queryable
	}

	rows, err := queryable.Query(selectQuery())
	if err != nil {
		return nil, err
	}

	records := []*Tax{}
	for rows.Next() {
		tax := &Tax{}
		err := scanTax(rows, tax)
		if err != nil {
			return nil, err
		}
		records = append(records, tax)
	}

	return records, nil
}

// FindByQuery ...
func (s *Storage) FindByQuery(ctx context.Context, query string) ([]*Tax, error) {
	queryable, ok := data.QueryableFromContext(ctx)
	if !ok {
		queryable = s.queryable
	}

	q := fmt.Sprintf(`%s %s`, selectQuery(), query)
	rows, err := queryable.Query(q)
	if err != nil {
		return nil, err
	}

	records := []*Tax{}
	for rows.Next() {
		tax := &Tax{}
		err := scanTax(rows, tax)
		if err != nil {
			return nil, err
		}
		records = append(records, tax)
	}

	return records, nil
}

// FindByKeys ...
func (s *Storage) FindByKeys(ctx context.Context, taxID int) (*Tax, error) {
	queryable, ok := data.QueryableFromContext(ctx)
	if !ok {
		queryable = s.queryable
	}

	query := fmt.Sprintf(`%s %s %v`, selectQuery(), `WHERE "taxId" =`, taxID)
	row := queryable.QueryRow(query)

	record := &Tax{}
	err := scanTax(row, record)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return record, nil
}

// Update ...
func (s *Storage) Update(ctx context.Context, tax *Tax) error {
	queryable, ok := data.QueryableFromContext(ctx)
	if !ok {
		queryable = s.queryable
	}

	query, params := updateQuery(tax)
	row := queryable.QueryRow(query, params...)

	err := scanTax(row, tax)
	if err != nil {
		return err
	}

	return nil
}
