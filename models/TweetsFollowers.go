package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*TweetsFollowers es la estructura con la que devolveremos los tweets */
type TweetsFollowers struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID     string             `bson:"UserID" json:"UserID,omitempty"`
	FollowerID string             `bson:"FollowerID" json:"FollowerID,omitempty"`
	Tweets     Tweet
}
