package tax

import (
	"context"

	"github.com/davidyunus/tax-calculator/internal/data"
)

// Service ...
type Service struct {
	queryable data.Queryable
	storage   *Storage
}

// NewService ...
func NewService(queryable data.Queryable) *Service {
	storage := NewStorage(queryable)
	return &Service{
		queryable: queryable,
		storage:   storage,
	}
}

// Create ...
func (s *Service) Create(ctx context.Context, tax *Tax) error {
	return s.storage.Create(ctx, tax)
}

// FindAll ...
func (s *Service) FindAll(ctx context.Context) ([]*Tax, error) {
	return s.storage.FindAll(ctx)
}

// FindByQuery ...
func (s *Service) FindByQuery(ctx context.Context, query string) ([]*Tax, error) {
	return s.storage.FindByQuery(ctx, query)
}

// FindByKeys ...
func (s *Service) FindByKeys(ctx context.Context, taxID int) (*Tax, error) {
	return s.storage.FindByKeys(ctx, taxID)
}

// Update ...
func (s *Service) Update(ctx context.Context, tax *Tax) error {
	return s.storage.Update(ctx, tax)
}
