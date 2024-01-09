package db

import (
	"context"
	"fmt"
	"go-Expense/pkg/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB(cfg config.Config) (*mongo.Client, error) {
	mongoUrl := fmt.Sprintf("host=%s dbname=%s password=%s port=%s", cfg.DB_HOST, cfg.DB_NAME, cfg.DB_PASSWORD, cfg.DB_PORT)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatal(err)
	}
	return client, nil
}

