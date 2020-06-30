package bd

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/jeananel/social.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*InsertTweet graba el Tweet en la BD */
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

/*GetTweets lee los tweets de un perfil */
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
