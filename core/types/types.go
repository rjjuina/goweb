package types

import "time"

type DBUser struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Age       int64     `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
