package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Christianesk/social-network-golang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRecord(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("social-network-golang")
	users := db.Collection("users")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := users.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	prueba := "asdasdas"
	fmt.Println(prueba)

	ObjectID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjectID.String(), true, nil

}
