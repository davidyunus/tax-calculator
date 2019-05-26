package data

import "database/sql"

// Manager ...
type Manager struct {
	db *sql.DB
}

// NewManager ...
func NewManager(db *sql.DB) *Manager {
	return &Manager{
		db: db,
	}
}
