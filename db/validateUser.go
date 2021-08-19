package db

import (
	"context"
	"time"

	"github.com/Christianesk/social-network-golang/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ValidateUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("social-network-golang")
	users := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := users.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID

}
