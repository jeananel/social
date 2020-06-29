package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoConnection variable for directly connect
var MongoConnection = ConnectionDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://sa:root@clustersocialnetwork-qba9d.mongodb.net/socialnetwork")

/*ConnectionDB connect to database*/
func ConnectionDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Printf("Success connection.")

	return client

}

/*CheckConnection verify connection*/
func CheckConnection() bool {
	err := MongoConnection.Ping(context.TODO(), nil)

	if err != nil {
		return false
	}

	return true

}
