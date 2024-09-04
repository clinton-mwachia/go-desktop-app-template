package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	Timestamp time.Time          `bson:"timestamp"`
	Details   string             `bson:"details,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
