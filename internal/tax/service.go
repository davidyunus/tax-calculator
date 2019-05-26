package tax

import (
	"context"
	"time"

	"github.com/tax-calculator/internal/data/tax"

	"github.com/tax-calculator/internal/data"
)

// Service ...
type Service struct {
	queryable  data.Queryable
	taxService *tax.Service
}

// NewService ...
func NewService(taxService *tax.Service) *Service {
	return &Service{
		taxService: taxService,
	}
}

// makeResponseTax create response of tax
func makeResponseTax(taxRequest *tax.Tax) *Tax {
	responseTax := &Tax{}
	typ := ""
	refund := ""
	var taxAmount float32

	switch taxRequest.TaxCode {
	case 1:
		typ = "food and beverage"
		refund = "yes"
		taxAmount = taxRequest.Price * 10 / 100
	case 2:
		typ = "tobacco"
		refund = "no"
		taxAmount = 10 + (0.02 * taxRequest.Price)
	case 3:
		typ = "entertainment"
		refund = "no"
		if taxRequest.Price >= 100 {
			t := taxRequest.Price - 100
			taxAmount = 0.01 * t
		}
	}

	responseTax.TaxID = taxRequest.TaxID
	responseTax.Name = taxRequest.Name
	responseTax.TaxCode = taxRequest.TaxCode
	responseTax.Type = typ
	responseTax.Refundable = refund
	responseTax.Price = taxRequest.Price
	responseTax.Tax = taxAmount
	responseTax.Amount = float32(taxAmount) + float32(taxRequest.Price)
	responseTax.CreatedAt = taxRequest.CreatedAt.Format(time.RFC3339)
	responseTax.UpdatedAt = taxRequest.UpdatedAt.Format(time.RFC3339)
	responseTax.DeletedAt = taxRequest.DeletedAt

	return responseTax
}

// Create ...
func (s *Service) Create(ctx context.Context, name string, taxCode int, price int) error {
	tx := &tax.Tax{}

	tx.Name = name
	tx.TaxCode = taxCode
	tx.Price = float32(price)
	utcNow := time.Now().UTC()
	tx.CreatedAt = &utcNow
	tx.UpdatedAt = &utcNow

	err := s.taxService.Create(ctx, tx)
	if err != nil {
		return err
	}
	return nil
}

// FindAll ...
func (s *Service) FindAll(ctx context.Context) ([]*Tax, error) {
	tx, err := s.taxService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	responseTax := []*Tax{}
	for i := range tx {
		responseTax = append(responseTax, makeResponseTax(tx[i]))
	}

	return responseTax, nil
}

// FindByQuery ...
func (s *Service) FindByQuery(ctx context.Context, query string) ([]*Tax, error) {
	whereQuery := `WHERE "name" LIKE '%` + query + `%' OR "taxCode" like '%` + query + `%' `

	tx, err := s.taxService.FindByQuery(ctx, whereQuery)
	if err != nil {
		return nil, err
	}

	responseTax := []*Tax{}
	for i := range tx {
		responseTax = append(responseTax, makeResponseTax(tx[i]))
	}

	return responseTax, nil
}

// FindByKeys ...
func (s *Service) FindByKeys(ctx context.Context, taxID int) (*Tax, error) {
	tx, err := s.taxService.FindByKeys(ctx, taxID)
	if err != nil {
		return nil, err
	}

	return makeResponseTax(tx), nil
}

// Update ...
func (s *Service) Update(ctx context.Context, taxID int, name string, taxCode, price int) error {
	tx, err := s.taxService.FindByKeys(ctx, taxID)
	if err != nil {
		return err
	}

	taxRequest := &tax.Tax{}
	taxRequest.TaxID = tx.TaxID
	taxRequest.Name = name
	taxRequest.TaxCode = taxCode
	taxRequest.Price = float32(price)
	utcNow := time.Now().UTC()
	taxRequest.UpdatedAt = &utcNow

	err = s.taxService.Update(ctx, taxRequest)
	if err != nil {
		return err
	}

	return nil
}
