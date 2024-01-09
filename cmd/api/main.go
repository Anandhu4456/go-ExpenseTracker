package main

import (
	"go-Expense/pkg/config"
	"log"
)


func main(){
	config,configErr:=config.LoadConfig()
	if configErr!=nil{
		log.Fatal("cannot load config ",configErr)
	}
	
}