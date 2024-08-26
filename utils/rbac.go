package utils

import (
	"context"
	"desktop-app-template/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ROLE BASED ACCESS CONTROL

// Check if the current user is an admin
func IsAdmin(userID primitive.ObjectID, window fyne.Window) bool {
	collection := GetCollection("users")
	var user models.User

	err := collection.FindOne(context.Background(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		dialog.ShowError(err, window)
		return false
	}

	return user.Role == "admin"
}
