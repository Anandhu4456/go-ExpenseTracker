package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct{
	ID primitive.ObjectID `bson:"id"`
	UserName string `json:"user_name"`
}