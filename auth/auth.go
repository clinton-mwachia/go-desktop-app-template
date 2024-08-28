package auth

import (
	"context"
	"desktop-app-template/models"
	"desktop-app-template/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(username, password, role string) error {
	if err := utils.ValidateUsername(username); err != nil {
		return err
	}

	if err := utils.ValidatePassword(password); err != nil {
		return err
	}

	if err := utils.ValidateRole(role); err != nil {
		return err
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	userCollection := utils.GetCollection("users")
	_, err = userCollection.InsertOne(context.Background(), models.User{
		ID:       primitive.NewObjectID(), // Generate a new ID for the user
		Username: username,
		Password: hashedPassword,
		Role:     role,
	})

	return err
}

func Login(username, password string) (*models.User, error) {
	userCollection := utils.GetCollection("users")
	var user models.User

	err := userCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	if !CheckPasswordHash(password, user.Password) {
		return nil, mongo.ErrNoDocuments
	}

	return &user, nil
}

// UpdateUserPassword updates the user's password in the database.
func UpdateUserPassword(userID primitive.ObjectID, password string, window fyne.Window) error {
	collection := utils.GetCollection("users")

	newHashedPassword, err := HashPassword(password)

	if err != nil {
		return err
	}

	// Update the user's password field in the database.
	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{"_id": userID},
		bson.M{"$set": bson.M{"password": newHashedPassword}},
	)

	if err != nil {
		dialog.ShowError(err, window)
		return err
	}

	return nil
}
