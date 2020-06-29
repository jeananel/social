package bd

import (
	"context"
	"time"

	"github.com/jeananel/social.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertRegister for register account
func InsertRegister(object models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	//When end instruction remove timeout operation and liberate context
	defer cancel()

	db := MongoConnection.Database("socialnetwork")
	collection := db.Collection("Users")

	//Set password encrypted
	passWordEncrypted, _ := EcryptPass(object.Password)
	object.Password = passWordEncrypted

	//passWordEncrypted2, _ := EcryptPasswordUtil(object.Password)

	result, err := collection.InsertOne(ctx, object)

	if err != nil {
		return "", false, err
	}

	//Get id of created object
	ObjectID, _ := result.InsertedID.(primitive.ObjectID)

	//Return created object id
	return ObjectID.String(), true, nil

}

//CheckExistUser for check user in database
func CheckExistUser(email string) (models.User, bool, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	//When end instruction remove timeout operation and liberate context
	defer cancel()

	db := MongoConnection.Database("socialnetwork")
	collection := db.Collection("Users")

	object := bson.M{"Email": email}

	var result models.User

	err := collection.FindOne(ctx, object).Decode(&result)

	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID, err
	}

	return result, true, ID, err

}
