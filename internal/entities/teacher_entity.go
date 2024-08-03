package entities

import "time"

type Teacher struct {
	ID         string    `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Department string    `json:"department"`
	DOB        string    `json:"dob"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
