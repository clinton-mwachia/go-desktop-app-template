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

// AddTodo adds a new todo to the database.
func AddTodo(todo models.Todo, window fyne.Window) {
	collection := GetCollection("todos")
	_, err := collection.InsertOne(context.TODO(), todo)
	if err != nil {
		dialog.ShowError(fmt.Errorf("failed to add todo: %v", err), window)
	} else {
		dialog.ShowInformation("Success", "Todo added successfully!", window)
	}
}

// BulkInsertTodos inserts multiple todos into the database.
func BulkInsertTodos(todos []models.Todo, window fyne.Window) {
	collection := GetCollection("todos")
	var documents []interface{}
	for _, todo := range todos {
		documents = append(documents, todo)
	}

	_, err := collection.InsertMany(context.TODO(), documents)
	if err != nil {
		dialog.ShowError(fmt.Errorf("failed to insert todos: %v", err), window)
	} else {
		dialog.ShowInformation("success", "todos inserted successfully!", window)
	}
}

// GetAllTodos retrieves all todos from the database.
func GetAllTodos(window fyne.Window) []models.Todo {
	collection := GetCollection("todos")
	var todos []models.Todo

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		dialog.ShowError(fmt.Errorf("failed to find todos: %v", err), window)
		return nil
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &todos); err != nil {
		dialog.ShowError(fmt.Errorf("failed to decode todos: %v", err), window)
	}

	return todos
}

// GetTodoByID retrieves a todo by its ID from the database.
func GetTodoByID(id primitive.ObjectID, window fyne.Window) *models.Todo {
	collection := GetCollection("todos")
	var todo models.Todo

	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		dialog.ShowError(fmt.Errorf("failed to find todo by ID: %v", err), window)
		return nil
	}

	return &todo
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
		dialog.ShowError(fmt.Errorf("failed to update todo: %v", err), window)
	} else {
		dialog.ShowInformation("success", "todo updated successfully!", window)
	}
}

// DeleteTodo deletes a todo from the database.
func DeleteTodo(id primitive.ObjectID, window fyne.Window) {
	collection := GetCollection("todos")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		dialog.ShowError(fmt.Errorf("failed to delete todo: %v", err), window)
	} else {
		dialog.ShowInformation("success", "todo deleted successfully!", window)
	}
}
