// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package store

import (
	"time"
)

type Source struct {
	ID        int32
	Name      string
	Url       string
	Driver    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
