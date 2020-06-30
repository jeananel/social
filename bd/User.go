package bd

import (
	"context"
	"time"

	"github.com/jeananel/social.git/models"
	"github.com/jeananel/social.git/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

/*GetUsers Get users  relations followers, following */
func GetUsers(ID string, page int64, search string, tipo string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("socialnetwork")
	collection := db.Collection("Users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"Name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := collection.Find(ctx, query, findOptions)
	if err != nil {
		return results, false
	}

	var found, flagInclude bool

	for cursor.Next(ctx) {
		var userObj models.User
		err := cursor.Decode(&userObj)
		if err != nil {
			return results, false
		}

		var userFollowerObj models.UsersFollowers
		userFollowerObj.UserID = ID

		userFollowerObj.FollowerID = userObj.ID.Hex()

		flagInclude = false

		found, err = CheckFollowing(userFollowerObj)

		//New users without relations. Users not following
		if tipo == "new" && found == false {
			flagInclude = true
		}

		//User following
		if tipo == "follow" && found == true {
			flagInclude = true
		}

		//Validation IDs not sames
		if userFollowerObj.FollowerID == ID {
			flagInclude = false
		}

		if flagInclude == true {
			//Clean values that not used
			userObj.Password = ""
			userObj.Biography = ""
			userObj.WebSite = ""
			userObj.Location = ""
			userObj.Banner = ""
			userObj.Email = ""

			results = append(results, &userObj)
		}
	}

	err = cursor.Err()
	if err != nil {
		return results, false
	}
	cursor.Close(ctx)
	return results, true
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

//UpdateUser for update user profile
func UpdateUser(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("socialnetwork")
	collection := db.Collection("Users")

	object := make(map[string]interface{})
	if len(u.Name) > 0 {
		object["Name"] = u.Name
	}
	if len(u.LastName) > 0 {
		object["LastName"] = u.LastName
	}
	object["DateBirth"] = u.DateBirth
	if len(u.Avatar) > 0 {
		object["Avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		object["Banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		object["Biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		object["Location"] = u.Location
	}
	if len(u.WebSite) > 0 {
		object["WebSite"] = u.WebSite
	}

	updtString := bson.M{
		"$set": object,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := collection.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}
