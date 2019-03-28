package main

import (
	"context"
	"fmt"
	"log"

//	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
//	"github.com/mongodb/mongo-go-driver/mongo/options"
)

type User struct {
	Name string
	Age  int
}

func main() {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

//	collection := client.Database("tutorial").Collection("users")
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
