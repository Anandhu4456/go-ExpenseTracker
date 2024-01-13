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
	mongoUrl := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatal("client not found: ", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("failed to connect : ", err)
	} else {
		log.Println("connected to mongodb...")
	}
	return client, nil
}
