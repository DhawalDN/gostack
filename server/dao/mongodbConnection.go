package dao

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Package DB Variable
var db *mongo.Database

//InitDB : Initialize mongoDB
func InitDB(url string, port int) error {

	//Using mongo library
	clientOptions := options.Client().ApplyURI("mongodb://" + url + ":" + strconv.Itoa(port))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("InitDB : ERROR_WHILE_CONNECTION")
		return errors.New("InitDB : ERROR_WHILE_CONNECTION")
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("InitDB : ERROR_WHILE_PINGING")
		return errors.New("InitDB : ERROR_WHILE_PINGING")
	}
	fmt.Println("Connected to mongo db")
	db = client.Database("gostack")
	return nil
}
