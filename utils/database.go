package utils

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB(uri string, window fyne.Window) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		dialog.ShowInformation("MongoDB Connect", "Failed to connect to MongoDB", window)
	}
	Client = client

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		dialog.ShowInformation("MongoDB Connect", "Failed to Ping MongoDB", window)
	}
}

func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("go_template_app").Collection(collectionName)
}
