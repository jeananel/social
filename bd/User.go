package bd

import (
	"context"
	"time"

	"github.com/jeananel/social.git/models"
	"github.com/jeananel/social.git/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

//InsertRegister for register account
func InsertRegister(object models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	//When end instruction remove timeout operation and liberate context
	defer cancel()

	db := MongoConnection.Database("socialnetwork")
	collection := db.Collection("Users")

	//Set password encrypted
	passWordEncrypted, _ := utils.EcryptPasswordUtil(object.Password)
	object.Password = passWordEncrypted

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
func CheckExistUser(email string) (models.User, bool, string) {
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
		return result, false, ID
	}

	return result, true, ID

}

//TryLoginUser for check usuario login
func TryLoginUser(email string, password string) (models.User, bool) {

	user, founded, _ := CheckExistUser(email)

	if !founded {
		return user, false
	}

	passwordBytes := []byte(password)     //Password login
	passwordUser := []byte(user.Password) //Password user

	err := bcrypt.CompareHashAndPassword(passwordUser, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true
}
