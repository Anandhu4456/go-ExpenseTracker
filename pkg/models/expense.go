package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Section struct {
	ID primitive.ObjectID `bson:"id"`
	UserName string `json:"user_name"`
	Buget float64 `json:"buget"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
}
