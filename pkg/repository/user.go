package repository

import (
	"context"
	"go-Expense/pkg/models"
	"go-Expense/pkg/repository/interfaces"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// constructor function

func NewUserRepo(client *mongo.Client, collection *mongo.Collection) interfaces.UserRepo {
	return UserRepo{
		client:     client,
		collection: collection,
	}
}

func (ur UserRepo) AddSection(userName string, buget float64, description string, createdAt time.Time) (models.Section, error) {

	coll := ur.client.Database("Expense").Collection(ur.collection.Name())
	if coll == nil {
		log.Fatal("failed to obtain collection",ur.collection.Name())
	}
	

	section := models.Section{
		UserName:    userName,
		Buget:       buget,
		Description: description,
		CreatedAt:   createdAt,
	}
	result, err := ur.collection.InsertOne(context.TODO(), section)
	if err != nil {
		return models.Section{}, err
	}
	resultId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Fatal("failed to get inserted id", err)
	}
	section.ID = resultId

	return section, nil

}
