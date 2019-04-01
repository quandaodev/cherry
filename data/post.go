package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type Post struct {
	Id        string
	Title     string
	ImageURL  string
	Content   string
	CreatedAt time.Time
}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// Get all posts in the database and returns it
func Posts() (posts []Post, err error) {
	log.Println("data Posts()")
	client := getDBClient()
	defer client.Disconnect(context.TODO())

	collection := client.Database("cherry").Collection("posts")
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal("Find Err:", err)
	}

	for cur.Next(context.TODO()) {

		elem := &bson.D{}
		if err = cur.Decode(elem); err != nil {
			log.Fatal("read Post: couldn't make post ready for display:", err)
		}
		m := elem.Map()
		p := Post{
			Id:       m["_id"].(primitive.ObjectID).Hex(),
			Title:    m["Title"].(string),
			ImageURL: m["ImageURL"].(string),
			Content:  m["Content"].(string),
			//CreatedAt: dateTimeMillis(m["createdAt"].(primitive.DateTime)),
		}
		posts = append(posts, p)
	}

	if err := cur.Err(); err != nil {
		log.Fatal("End Err:", err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found posts (array of pointers): %+v\n", posts)
	return
}
