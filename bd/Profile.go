package bd

import (
	"context"
	"time"

	"github.com/jeananel/social.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*FindProfile find profile by id in database */
func FindProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConnection.Database("socialnetwork")
	collection := db.Collection("Users")

	var result models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	object := bson.M{
		"_id": objID,
	}

	err := collection.FindOne(ctx, object).Decode(&result)
	result.Password = ""
	if err != nil {
		return result, err
	}
	return result, nil
}
