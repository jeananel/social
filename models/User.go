package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User for main information about user
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"Name" bson:"Name"`
	LastName  string             `json:"LastName" bson:"LastName"`
	DateBirth time.Time          `json:"DateBirth" bson:"DateBirth"`
	Email     string             `json:"Email" bson:"Email"`
	Password  string             `json:"Password" bson:"Password"`
	Avatar    string             `json:"Avatar" bson:"Avatar,omitempty"`
	Banner    string             `json:"Banner" bson:"Banner,omitempty"`
	Biography string             `json:"Biography" bson:"Biography,omitempty"`
	Location  string             `json:"Location" bson:"Location,omitempty"`
	WebSite   string             `json:"WebSite" bson:"WebSite,omitempty"`
}
