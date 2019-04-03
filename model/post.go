package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post Struct
type Post struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"Title,omitempty" bson:"Title,omitempty"`
	ImageURL string             `json:"ImageURL,omitempty" bson:"ImageURL,omitempty"`
	Content  string             `json:"Content,omitempty" bson:"Content,omitempty"`
	CreateAt time.Time
}

/*
// Create at date
func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}*/

// Get all posts in the database and returns it
func ListPosts() (posts []Post, err error) {
	log.Println("listPosts() called")
	client := getDBClient()
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	defer client.Disconnect(ctx)
	collection := client.Database("cherry").Collection("posts")
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal("Find Err:", err)
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var post Post
		cur.Decode(&post)
		posts = append(posts, post)
	}
	fmt.Printf("Found posts (array of pointers): %+v\n", posts)
	return
}

// GetPostByID return a post matching the postID
func GetPostByID(postID string) (post Post, err error) {
	log.Println("getPostById(", postID, ") called")
	client := getDBClient()
	defer client.Disconnect(context.TODO())

	collection := client.Database("cherry").Collection("posts")
	id, _ := primitive.ObjectIDFromHex(postID)
	collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&post)
	fmt.Println("Found post with ID ", postID, post)
	return
}
