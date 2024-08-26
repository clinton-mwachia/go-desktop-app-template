package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	Message   string             `bson:"message"`
	CreatedAt time.Time          `bson:"created_at"`
	Read      bool               `bson:"read"`
}
