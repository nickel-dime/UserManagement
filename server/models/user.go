package models

import "time"

// creates a user
type User struct {
	ID           string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	RegisteredAt time.Time `json:"registeredAt,omitempty"`
}
