// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID            int64          `json:"id"`
	CreatedBy     sql.NullString `json:"created_by"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedBy     sql.NullString `json:"updated_by"`
	UpdatedAt     time.Time      `json:"updated_at"`
	MarkForDelete bool           `json:"mark_for_delete"`
	Owner         string         `json:"owner"`
	Balance       int64          `json:"balance"`
	Currency      string         `json:"currency"`
}

type BaseEntity struct {
	ID            int64          `json:"id"`
	CreatedBy     sql.NullString `json:"created_by"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedBy     sql.NullString `json:"updated_by"`
	UpdatedAt     time.Time      `json:"updated_at"`
	MarkForDelete bool           `json:"mark_for_delete"`
}

type Entry struct {
	ID            int64          `json:"id"`
	CreatedBy     sql.NullString `json:"created_by"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedBy     sql.NullString `json:"updated_by"`
	UpdatedAt     time.Time      `json:"updated_at"`
	MarkForDelete bool           `json:"mark_for_delete"`
	AccountID     sql.NullInt64  `json:"account_id"`
	// can be negative or positive
	Amount int64 `json:"amount"`
}

type Transfer struct {
	ID            int64          `json:"id"`
	CreatedBy     sql.NullString `json:"created_by"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedBy     sql.NullString `json:"updated_by"`
	UpdatedAt     time.Time      `json:"updated_at"`
	MarkForDelete bool           `json:"mark_for_delete"`
	FromAccountID sql.NullInt64  `json:"from_account_id"`
	ToAccountID   sql.NullInt64  `json:"to_account_id"`
	// mmust be positive
	Amount int64 `json:"amount"`
}
