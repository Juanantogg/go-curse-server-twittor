package db

import (
	"context"
	"log"

	"github.com/Juanantogg/server-go-twittor/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbUser = utils.GoDotEnvVariable("DB_USER")
var dbPassword = utils.GoDotEnvVariable("DB_PASSWORD")
var dbURI = utils.GoDotEnvVariable("DB_URI")

var MongoConnection = DBConnect()
var clientOptions = options.Client().ApplyURI("mongodb+srv://" + dbUser + ":" + dbPassword + dbURI)

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
