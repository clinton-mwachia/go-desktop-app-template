package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    UserID   primitive.ObjectID `bson:"user_id"`
    Title    string             `bson:"title"`
    Content  string             `bson:"content"`
    Done     bool               `bson:"done"`
}
