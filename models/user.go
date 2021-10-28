package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	FirstName string             `bson:"firstName" json:"firstName,omitempty"`
	LastName  string             `bson:"lastName" json:"lastName,omitempty"`
	Birthday  time.Time          `bson:"birthday" json:"birthday,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"Banner" json:"Banner,omitempty"`
	Biography string             `bson:"biography" json:"biography,omitempty"`
	Ubication string             `bson:"ubication" json:"ubication,omitempty"`
	Website   string             `bson:"website" json:"website,omitempty"`
}
