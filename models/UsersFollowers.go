package models

/*UsersFollowers relations between users */
type UsersFollowers struct {
	UserID     string `bson:"UserID" json:"UserID"`
	FollowerID string `bson:"FollowerID" json:"FollowerID"`
}
