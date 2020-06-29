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
	Avatar    string             `json:"Avatar" bson:"Avatar"`
	Banner    string             `json:"Banner" bson:"Banner"`
	Biography string             `json:"Biography" bson:"Biography"`
	Location  string             `json:"Location" bson:"Location"`
	WebSite   string             `json:"WebSite" bson:"WebSite"`
}
