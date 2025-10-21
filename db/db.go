package db

import (
	"context"
	"log"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client
var dbName string

func Connect() {
	_client, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	dbName = os.Getenv("MONGO_DB")

	if err != nil {
		log.Fatalln("Failed to connect to database", err)
	}
	err = _client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalln("Failed to ping database", err)
	}
	defer Disconnect()
	log.Println("Connected to database")
	client = _client
}

func Disconnect() {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatalln("Failed to disconnect gracefully", err)
	}
	log.Println("Disconnected from database")
}

func GetServer(name string) (*Server, error) {
	collection := GetServerCollection()
	var result Server
	err := collection.FindOne(context.TODO(), bson.D{{Key: "name", Value: strings.ToLower(name)}}).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func UpdateServer(server *Server) error {
	collection := GetServerCollection()

	server.Name = strings.ToLower(server.Name)
	_, err := GetServer(server.Name)

	if err != nil {
		_, err := collection.InsertOne(context.TODO(), server)

		if err != nil {
			return err
		}
	}

	_, err = collection.ReplaceOne(context.TODO(), bson.D{{Key: "name", Value: server.Name}}, server)

	if err != nil {
		return err
	}

	return nil
}

func GetServerCollection() *mongo.Collection {
	return client.Database(dbName).Collection("servers")
}
