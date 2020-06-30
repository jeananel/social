package bd

import (
	"context"
	"log"
	"time"

	"github.com/jeananel/social.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*InsertTweet save user tweet in database */
func InsertTweet(t models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("socialnetwork")
	collection := db.Collection("Tweets")

	object := bson.M{
		"UserID":  t.UserID,
		"Message": t.Message,
		"Date":    t.Date,
	}
	result, err := collection.InsertOne(ctx, object)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}

/*GetTweets get tweets of user in session */
func GetTweets(ID string, pagina int64) ([]*models.Tweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConnection.Database("socialnetwork")
	collection := db.Collection("Tweets")

	var results []*models.Tweet

	filter := bson.M{
		"UserID": ID,
	}

	configOptions := options.Find()
	configOptions.SetLimit(20)
	configOptions.SetSort(bson.D{{Key: "Date", Value: -1}})
	configOptions.SetSkip((pagina - 1) * 20)

	cursor, err := collection.Find(ctx, filter, configOptions)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {

		var registro models.Tweet
		err := cursor.Decode(&registro)
		if err != nil {
			log.Fatal(err.Error())
			return results, false
		}
		results = append(results, &registro)
	}
	return results, true
}

/*DeleteTweet delete tweet by ID */
func DeleteTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConnection.Database("socialnetwork")
	collection := db.Collection("Tweets")

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{
		"_id":    objID,
		"UserID": UserID,
	}

	_, err := collection.DeleteOne(ctx, filter)
	return err
}

/*GetTweetsFollowes tweets of followers */
func GetTweetsFollowes(ID string, pagina int) ([]models.TweetsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("socialnetwork")
	collection := db.Collection("UsersFollowers")

	skip := (pagina - 1) * 20

	//Math between tables
	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"UserID": ID}})

	//Unions tables relations
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "Tweets",     //Table
			"localField":   "FollowerID", //Relation
			"foreignField": "UserID",     //Relation
			"as":           "Tweets",     //Alias
		}})

	//List of one level -- unwind
	conditions = append(conditions, bson.M{"$unwind": "$Tweets"})

	//Filters
	conditions = append(conditions, bson.M{"$sort": bson.M{"Tweets.Date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := collection.Aggregate(ctx, conditions)
	var result []models.TweetsFollowers
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
