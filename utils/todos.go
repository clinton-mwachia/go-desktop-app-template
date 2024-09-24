package utils

import (
	"context"
	"desktop-app-template/models"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AddTodo adds a new todo to the database.
func AddTodo(todo models.Todo, window fyne.Window) {
	collection := GetCollection("todos")
	_, err := collection.InsertOne(context.TODO(), todo)
	if err != nil {
		dialog.ShowError(err, window)
	} else {
		dialog.ShowInformation("Success", "Todo added", window)
	}
}

// GetAllTodos retrieves all todos from the database.
func GetAllTodos(window fyne.Window) []models.Todo {
	collection := GetCollection("todos")
	var todos []models.Todo

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		dialog.ShowError(err, window)
		return todos
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &todos); err != nil {
		dialog.ShowError(err, window)
	}

	return todos
}

// GetTodoByID retrieves a single todo by its ID from the database.
func GetTodoByID(id primitive.ObjectID, window fyne.Window) models.Todo {
	collection := GetCollection("todos")
	var todo models.Todo

	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		dialog.ShowError(err, window)
	}

	return todo
}

// GetTodosByUserID retrieves all todos associated with a specific user.
func GetTodosByUserID(userID primitive.ObjectID, window fyne.Window) []models.Todo {
	collection := GetCollection("todos")
	var todos []models.Todo

	cursor, err := collection.Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		dialog.ShowError(err, window)
		return todos
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &todos); err != nil {
		dialog.ShowError(err, window)
	}

	return todos
}

// BulkInsertTodos inserts multiple todos into the database.
func BulkInsertTodos(todos []models.Todo, userID primitive.ObjectID, window fyne.Window, progressBar *widget.ProgressBar) {
	collection := GetCollection("todos")
	var docs []interface{}
	totalTodos := len(todos)
	progress := 0

	for i, todo := range todos {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		todo.CreatedAt = parsedTime
		todo.UserID = userID
		docs = append(docs, todo)

		// Update progress bar for each todo processed
		progress = i + 1
		progressBar.SetValue(float64(progress) / float64(totalTodos))

		// Flush the documents in smaller batches
		if len(docs) == 100 || i == totalTodos-1 {
			_, err := collection.InsertMany(context.TODO(), docs)
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			docs = nil // Reset docs slice for next batch
		}
	}

	dialog.ShowInformation("Success", "Todos added successfully!", window)
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
	} else {
		dialog.ShowInformation("Success", "Todo updated successfully!", window)
	}
}

// DeleteTodo deletes a todo from the database.
func DeleteTodo(id primitive.ObjectID, window fyne.Window) {
	collection := GetCollection("todos")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		dialog.ShowError(err, window)
	} else {
		dialog.ShowInformation("Success", "Todo deleted successfully!", window)
	}
}

// GetTodosPaginated fetches todos with pagination from the database
func GetTodosPaginated(page, limit int, userID primitive.ObjectID, w fyne.Window) []models.Todo {
	collection := GetCollection("todos")

	skip := (page - 1) * limit
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))

	var todos []models.Todo

	cursor, err := collection.Find(context.TODO(), bson.M{"user_id": userID}, findOptions)
	if err != nil {
		dialog.ShowError(err, w)
		return todos
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &todos); err != nil {
		dialog.ShowError(err, w)
	}

	return todos
}

// CountTodos returns the total count of todos for a user
func CountTodos(userID primitive.ObjectID, w fyne.Window) int64 {
	collection := GetCollection("todos")
	count, err := collection.CountDocuments(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		dialog.ShowError(err, w)
	}
	return count
}

// search todos by quering the db
func SearchTodos(searchText string, userID primitive.ObjectID, window fyne.Window) []models.Todo {
	collection := GetCollection("todos")

	// Create a case-insensitive regex pattern for the search
	searchPattern := bson.M{
		"$regex":   searchText,
		"$options": "i", // Case-insensitive
	}

	filter := bson.M{
		"user_id": userID,
		"$or": []bson.M{
			{"title": searchPattern},
			{"content": searchPattern},
		},
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		dialog.ShowError(err, window)
		return nil
	}
	defer cursor.Close(context.TODO())

	var results []models.Todo
	if err = cursor.All(context.TODO(), &results); err != nil {
		dialog.ShowError(err, window)
		return nil
	}

	return results

}

// get todo statistics fromthe db
func FetchTodoStatistics(userID primitive.ObjectID, window fyne.Window) (int, int, int) {
	collection := GetCollection("todos")

	// Context with timeout to avoid long-running queries
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Total Todos count
	totalCount, err := collection.CountDocuments(ctx, bson.M{"user_id": userID})
	if err != nil {
		dialog.ShowError(err, window)
		return 0, 0, 0
	}

	// Completed Todos count (Done: true)
	completedCount, err := collection.CountDocuments(ctx, bson.M{"user_id": userID, "done": true})
	if err != nil {
		dialog.ShowError(err, window)
		return int(totalCount), 0, 0
	}

	// Pending Todos count (Done: false)
	pendingCount, err := collection.CountDocuments(ctx, bson.M{"user_id": userID, "done": false})
	if err != nil {
		dialog.ShowError(err, window)
		return int(totalCount), int(completedCount), 0
	}

	return int(totalCount), int(completedCount), int(pendingCount)
}

// FetchTodoDataForCharts fetches the todos data grouped by done status and creation month/year for charts.
func FetchTodoDataForCharts(userID primitive.ObjectID, window fyne.Window) (map[string]int, map[string]int) {
	collection := GetCollection("todos")

	// Context with timeout to avoid long-running queries
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Bar chart data: count of todos grouped by done bool
	doneData := map[string]int{"true": 0, "false": 0}
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "user_id", Value: userID}}}},
		{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$done"}, {Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}}}}},
	})
	if err != nil {
		dialog.ShowError(err, window)
		return doneData, nil
	}
	defer cursor.Close(ctx)

	// Process bar chart data
	for cursor.Next(ctx) {
		var result struct {
			ID    bool `bson:"_id"`
			Count int  `bson:"count"`
		}
		if err := cursor.Decode(&result); err != nil {
			dialog.ShowError(err, window)
			return doneData, nil
		}

		if result.ID {
			doneData["true"] = result.Count
		} else {
			doneData["false"] = result.Count
		}
	}

	// Check if there were any errors during the cursor iteration
	if err := cursor.Err(); err != nil {
		dialog.ShowError(err, window)
		return doneData, nil
	}

	// Line chart data: count of todos grouped by creation month/year
	dateData := make(map[string]int)
	cursor, err = collection.Aggregate(ctx, mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "user_id", Value: userID}}}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: bson.D{
				{Key: "year", Value: bson.D{{Key: "$year", Value: "$created_at"}}},
				{Key: "month", Value: bson.D{{Key: "$month", Value: "$created_at"}}},
			}},
			{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
		}}},
	})
	if err != nil {
		dialog.ShowError(err, window)
		return doneData, dateData
	}
	defer cursor.Close(ctx)

	// Process line chart data
	for cursor.Next(ctx) {
		var result struct {
			ID struct {
				Year  int `bson:"year"`
				Month int `bson:"month"`
			} `bson:"_id"`
			Count int `bson:"count"`
		}
		if err := cursor.Decode(&result); err != nil {
			dialog.ShowError(err, window)
			return doneData, dateData
		}

		// Format as "YYYY-MM" for grouping by month and year
		dateKey := fmt.Sprintf("%d-%02d", result.ID.Year, result.ID.Month)
		dateData[dateKey] = result.Count
	}

	// Check if there were any errors during the cursor iteration
	if err := cursor.Err(); err != nil {
		dialog.ShowError(err, window)
	}

	return doneData, dateData
}
