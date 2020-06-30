package bd

import (
	"context"
	"time"

	"github.com/jeananel/social.git/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*InsertUsersFollowers relations between users BD */
func InsertUsersFollowers(userFollower models.UsersFollowers) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("socialnetwork")
	collection := db.Collection("UsersFollowers")

	_, err := collection.InsertOne(ctx, userFollower)
	if err != nil {
		return false, err
	}

	return true, nil
}

/*DeleteUsersFollowers remove relations between users BD */
func DeleteUsersFollowers(userFollower models.UsersFollowers) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("socialnetwork")
	collection := db.Collection("UsersFollowers")

	_, err := collection.DeleteOne(ctx, userFollower)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*CheckFollowing consulta la relacion entre 2 usuarios */
func CheckFollowing(userFollower models.UsersFollowers) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("socialnetwork")
	collection := db.Collection("UsersFollowers")

	filter := bson.M{
		"UserID":     userFollower.UserID,
		"FollowerID": userFollower.FollowerID,
	}

	var result models.UsersFollowers
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return false, err
	}
	return true, nil
}
