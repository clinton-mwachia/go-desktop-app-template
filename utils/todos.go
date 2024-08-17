package utils

import (
	"context"
	"desktop-app-template/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddTodo adds a new todo to the database.
func AddTodo(todo models.Todo, window fyne.Window) {
	collection := GetCollection("todos")
	_, err := collection.InsertOne(context.TODO(), todo)
	if err != nil {
		dialog.ShowError(err, window)
		return
	}
	dialog.ShowInformation("Success", "Todo added successfully!", window)
}

// GetAllTodos retrieves all todos from the database.
func GetAllTodos(window fyne.Window) []models.Todo {
	collection := GetCollection("todos")
	var todos []models.Todo

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		dialog.ShowError(err, window)
		return nil
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &todos); err != nil {
		dialog.ShowError(err, window)
		return nil
	}

	return todos
}

// UpdateTodo updates an existing todo in the database.
func UpdateTodo(todo models.Todo, window fyne.Window) {
	collection := GetCollection("todos")
	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": todo.ID},
		bson.M{"$set": todo},
	)
	if err != nil {
		dialog.ShowError(err, window)
		return
	}
	dialog.ShowInformation("Success", "Todo updated successfully!", window)
}

// DeleteTodo deletes a todo from the database.
func DeleteTodo(id primitive.ObjectID, window fyne.Window) {
	collection := GetCollection("todos")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		dialog.ShowError(err, window)
		return
	}
	dialog.ShowInformation("Success", "Todo deleted successfully!", window)
}
