package auth

import (
	"context"
	"desktop-app-template/models"
	"desktop-app-template/utils"

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
