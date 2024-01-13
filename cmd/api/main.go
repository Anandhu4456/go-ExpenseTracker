package main

import (
	"go-Expense/pkg/config"
	"go-Expense/pkg/di"
	"log"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config ", configErr)
	}
	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatal("couldn't start server", err)
	} else {
		server.StartServer()
	}
}
