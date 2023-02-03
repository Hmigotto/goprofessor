package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StartMongoConnection() (*mongo.Client, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return nil, err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func ConectaComBancoDeDadosColectionPeople() *mongo.Collection {

	client, err := StartMongoConnection()

	if err != nil {
		return nil
	}

	db := client.Database("Golearn")
	people := db.Collection("people")
	return people
}
