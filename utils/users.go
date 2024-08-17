package utils

import (
	"context"
	"desktop-app-template/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddUser adds a new user to the database.
func AddUser(user models.User, window fyne.Window) {
	collection := GetCollection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		dialog.ShowError(err, window)
	} else {
		dialog.ShowInformation("Success", "User added successfully!", window)
	}
}

// GetAllUsers retrieves all users from the database.
func GetAllUsers(window fyne.Window) []models.User {
	collection := GetCollection("users")
	var users []models.User

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		dialog.ShowError(err, window)
		return users
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &users); err != nil {
		dialog.ShowError(err, window)
	}

	return users
}

// GetUserByID retrieves a single user by its ID from the database.
func GetUserByID(id primitive.ObjectID, window fyne.Window) models.User {
	collection := GetCollection("users")
	var user models.User

	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		dialog.ShowError(err, window)
	}

	return user
}

// UpdateUser updates an existing user in the database.
func UpdateUser(user models.User, window fyne.Window) {
	collection := GetCollection("users")
	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": user.ID},
		bson.M{"$set": user},
	)
	if err != nil {
		dialog.ShowError(err, window)
	} else {
		dialog.ShowInformation("Success", "User updated successfully!", window)
	}
}

// DeleteUser deletes a user from the database.
func DeleteUser(id primitive.ObjectID, window fyne.Window) {
	collection := GetCollection("users")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		dialog.ShowError(err, window)
	} else {
		dialog.ShowInformation("Success", "User deleted successfully!", window)
	}
}
