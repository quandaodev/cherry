package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
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

	collection := client.Database("tutorial").Collection("users")

	/*
		peter := User{"Peter", 10}
		insertResult, err := collection.InsertOne(context.TODO(), peter)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)

		tom := User{"Tom", 10}
		collection.InsertOne(context.TODO(), tom)
	*/
	// Pass these options to the Find method
	//findOptions := options.Find()
	//findOptions.SetLimit(2)

	// Here's an array in which you can store the decoded documents
	var results []*User

	// Passing nil as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{}) //collection.Find(context.TODO(), nil, findOptions)
	if err != nil {
		log.Fatal("Find Err:", err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal("Decode Err:", err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal("End Err:", err)
	}
	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	for i, v := range results {
		fmt.Println("User ", i, ":", v.Name, " ", v.Age)
	}

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal("Disconnect Err:", err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
