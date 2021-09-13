package db

import (
	"context"
	"time"

	"github.com/Christianesk/social-network-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateRecord(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("social-network-golang")
	users := db.Collection("users")

	record := make(map[string]interface{})

	if len(user.Name) > 0 {
		record["name"] = user.Name
	}
	if len(user.Lastname) > 0 {
		record["lastName"] = user.Lastname
	}
	record["birthdate"] = user.Birthdate
	if len(user.Email) > 0 {
		record["email"] = user.Email
	}
	if len(user.Password) > 0 {
		record["password"] = user.Password
	}
	if len(user.Avatar) > 0 {
		record["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		record["banner"] = user.Banner
	}
	if len(user.Biography) > 0 {
		record["biography"] = user.Biography
	}
	if len(user.Location) > 0 {
		record["location"] = user.Location
	}
	if len(user.Website) > 0 {
		record["website"] = user.Website
	}

	updateString := bson.M{
		"$set": record,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{
		"_id": bson.M{"$eq": objID},
	}

	_, err := users.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}

	return true, nil

}
