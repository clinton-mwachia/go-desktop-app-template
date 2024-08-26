package utils

import (
	"context"
	"desktop-app-template/models"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddNotification adds a new notification to the database
func AddNotification(notification models.Notification, window fyne.Window) {
	collection := GetCollection("notifications")
	notification.CreatedAt = time.Now()
	_, err := collection.InsertOne(context.TODO(), notification)
	if err != nil {
		dialog.ShowError(err, window)
	}
}

// GetNotifications retrieves all notifications for a user
func GetNotifications(userID primitive.ObjectID, window fyne.Window) []models.Notification {
	var notifications []models.Notification
	collection := GetCollection("notifications")
	filter := bson.M{"user_id": userID}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		dialog.ShowError(err, window)
		return notifications
	}

	if err = cursor.All(context.TODO(), &notifications); err != nil {
		dialog.ShowError(err, window)
	}
	return notifications
}

// MarkNotificationAsRead updates a notification's read status in the database
func MarkNotificationAsRead(notificationID primitive.ObjectID, window fyne.Window) {
	collection := GetCollection("notifications")
	filter := bson.M{"_id": notificationID}
	update := bson.M{"$set": bson.M{"read": true}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		dialog.ShowError(err, window)
	}
}

// ClearNotifications deletes all notifications for a user
func ClearNotifications(userID primitive.ObjectID, window fyne.Window) {
	collection := GetCollection("notifications")
	filter := bson.M{"user_id": userID}
	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		dialog.ShowError(err, window)
	}
}
