package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Christianesk/social-network-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("social-network-golang")
	users := db.Collection("users")

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{
		"_id": objID,
	}

	err := users.FindOne(ctx, condition).Decode(&profile)

	profile.Password = ""
	if err != nil {
		fmt.Println("Record not found" + err.Error())
		return profile, err
	}
	return profile, nil
}
