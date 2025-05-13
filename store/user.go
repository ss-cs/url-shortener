package store

import (
	"context"
	"errors"
	"url-shortener/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(email, password string) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user := models.User{Email: email, Password: string(hashed)}
	_, err := collectionUser.InsertOne(context.Background(), user)
	return err
}

func ValidateUser(email, password string) error {
	var user models.User
	err := collectionUser.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return errors.New("user not found")
	}
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
