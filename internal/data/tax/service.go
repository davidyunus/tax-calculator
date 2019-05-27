package tax

import (
	"context"

	"github.com/davidyunus/tax-calculator/internal/data"
)

// Service represent service tax data
type Service struct {
	queryable data.Queryable
	storage   *Storage
}

// NewService create new tax service
func NewService(queryable data.Queryable) *Service {
	storage := NewStorage(queryable)
	return &Service{
		queryable: queryable,
		storage:   storage,
	}
}

// Create , create tax data
func (s *Service) Create(ctx context.Context, tax *Tax) error {
	return s.storage.Create(ctx, tax)
}

// FindAll , find all taxes data
func (s *Service) FindAll(ctx context.Context) ([]*Tax, error) {
	return s.storage.FindAll(ctx)
}

// FindByQuery , find tax data by query
func (s *Service) FindByQuery(ctx context.Context, query string) ([]*Tax, error) {
	return s.storage.FindByQuery(ctx, query)
}

// FindByKeys , find tax data by taxId
func (s *Service) FindByKeys(ctx context.Context, taxID int) (*Tax, error) {
	return s.storage.FindByKeys(ctx, taxID)
}

// Update , update tax data
func (s *Service) Update(ctx context.Context, tax *Tax) error {
	return s.storage.Update(ctx, tax)
}
