package data

import (
	"context"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func getDBClient() (client *mongo.Client) {
	client, err := mongo.Connect(context.TODO(), "mongodb://quandao.me:27017")

	if err != nil {
		log.Fatal(err)
	}
	return
}

/*
func insertItem(client *mongo.Client) {
	collection := client.Database("tutorial").Collection("users")
	//peter := User{"Peter", 10}
	//insertResult, err := collection.InsertOne(context.TODO(), peter)
	//if err != nil {
	//log.Fatal(err)
	//}
}


func listItems(client *mongo.Client) {
	collection := client.Database("tutorial").Collection("users")

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
	//for i, v := range results {
	//	fmt.Println("User ", i, ":", v.Name, " ", v.Age)
	//}
}*/

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
