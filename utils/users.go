package utils

import (
	"context"
	"desktop-app-template/models"
	"fmt"

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
		dialog.ShowError(fmt.Errorf("Failed to add user: %v", err), window)
	} else {
		dialog.ShowInformation("Success", "User added successfully!", window)
	}
}

// BulkInsertUsers inserts multiple users into the database.
func BulkInsertUsers(users []models.User, window fyne.Window) {
	collection := GetCollection("users")
	var documents []interface{}
	for _, user := range users {
		documents = append(documents, user)
	}

	_, err := collection.InsertMany(context.TODO(), documents)
	if err != nil {
		dialog.ShowError(fmt.Errorf("Failed to insert users: %v", err), window)
	} else {
		dialog.ShowInformation("Success", "Users inserted successfully!", window)
	}
}

// GetAllUsers retrieves all users from the database.
func GetAllUsers(window fyne.Window) []models.User {
	collection := GetCollection("users")
	var users []models.User

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		dialog.ShowError(fmt.Errorf("Failed to find users: %v", err), window)
		return nil
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &users); err != nil {
		dialog.ShowError(fmt.Errorf("Failed to decode users: %v", err), window)
	}

	return users
}

// GetUserByID retrieves a user by its ID from the database.
func GetUserByID(id primitive.ObjectID, window fyne.Window) *models.User {
	collection := GetCollection("users")
	var user models.User

	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		dialog.ShowError(fmt.Errorf("Failed to find user by ID: %v", err), window)
		return nil
	}

	return &user
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
		dialog.ShowError(fmt.Errorf("Failed to update user: %v", err), window)
	} else {
		dialog.ShowInformation("Success", "User updated successfully!", window)
	}
}

// DeleteUser deletes a user from the database.
func DeleteUser(id primitive.ObjectID, window fyne.Window) {
	collection := GetCollection("users")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		dialog.ShowError(fmt.Errorf("Failed to delete user: %v", err), window)
	} else {
		dialog.ShowInformation("Success", "User deleted successfully!", window)
	}
}
