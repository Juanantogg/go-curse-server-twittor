package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnection = DBConnect()
var clientOptions = options.Client().ApplyURI("mongodb+srv://dbUser:dbUser@twittor-0.ykk4r.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

func DBConnect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())

		return client
	}

	// optional
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())

		return client
	}

	log.Println("successful database connection")

	return client
}

func ConnectionCheck() bool {
	err := MongoConnection.Ping(context.TODO(), nil)

	return err != nil
}
