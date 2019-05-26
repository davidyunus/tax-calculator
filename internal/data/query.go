package data

import (
	"context"
	"database/sql"
)

type key int

const queryableKey key = 0

// Queryable ...
type Queryable interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// NewContext ...
func NewContext(ctx context.Context, q Queryable) context.Context {
	return context.WithValue(ctx, queryableKey, q)
}

// QueryableFromContext ...
func QueryableFromContext(ctx context.Context) (Queryable, bool) {
	q, ok := ctx.Value(queryableKey).(Queryable)
	return q, ok
}
