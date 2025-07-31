package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Phone     string    `bson:"phone"`
	CreatedAt time.Time `bson:"created_at"`
}
