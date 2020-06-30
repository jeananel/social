package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Tweet estructure Tweet in database */
type Tweet struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"UserID" json:"UserID"`
	Message string             `bson:"Message" json:"Message"`
	Date    time.Time          `bson:"Date" json:"Date"`
}
