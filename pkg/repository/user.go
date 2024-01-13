package repository

import (
	"context"
	"go-Expense/pkg/models"
	"go-Expense/pkg/repository/interfaces"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	client *mongo.Client
	// collection *mongo.Collection
}

// constructor function

func NewUserRepo(client *mongo.Client /*collection *mongo.Collection*/) interfaces.UserRepo {
	return UserRepo{
		client: client,
		// collection: collection,
	}
}

func (ur UserRepo) AddSection(userName string, buget float64, description string) (models.Section, error) {

	coll := ur.client.Database("Expense").Collection("section")
	if coll == nil {
		log.Fatal("failed to obtain collection", coll)
	}

	section := models.Section{
		UserName:    userName,
		Buget:       buget,
		Description: description,
		// CreatedAt:   createdAt,
	}
	result, err := coll.InsertOne(context.TODO(), section)
	if err != nil {
		log.Println(err)
		return models.Section{}, err
	}
	log.Println("section added with id :", result.InsertedID)
	return section, nil
}
