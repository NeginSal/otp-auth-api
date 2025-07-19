package model

import (
	"time"
)

type User struct {
	ID        string    `bson:"_id,omitempty"`
	Phone     string    `bson:"phone"`
	CreatedAt time.Time `bson:"created_at"`
}
