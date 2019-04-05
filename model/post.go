package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	log.Println("ListPosts() called")
	client := getDBClient()
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	defer client.Disconnect(ctx)
	collection := client.Database("cherry").Collection("posts")
	options := options.FindOptions{}
	options.Sort = bson.D{{"_id", -1}}
	cur, err := collection.Find(context.TODO(), bson.D{}, &options)
	if err != nil {
		log.Fatal("Find Err:", err)
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var post Post
		cur.Decode(&post)
		posts = append(posts, post)
	}
	return
}

// GetPostByID return a post matching the postID
func GetPostByID(postID string) (post Post, err error) {
	log.Println("GetPostById(", postID, ") called")
	client := getDBClient()
	defer client.Disconnect(context.TODO())

	collection := client.Database("cherry").Collection("posts")
	id, _ := primitive.ObjectIDFromHex(postID)
	collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&post)
	return
}

// CreatePost inserts a new post to the database
func CreatePost(newPost Post) (err error) {
	log.Println("CreatePost called")
	client := getDBClient()
	defer client.Disconnect(context.TODO())

	collection := client.Database("cherry").Collection("posts")
	_, err = collection.InsertOne(context.TODO(), newPost)
	if err != nil {
		log.Fatal(err)
	}

	return
}
