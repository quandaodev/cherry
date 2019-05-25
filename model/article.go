package model

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

// Post Struct
type Article struct {
	Title   string `firestore:"title,omitempty"`
	Content string `firestore:"content,omitempty"`
	Slug    string `firestore:"slug,omitempty"`
}

// Get all posts in the database and returns it
func ListArticles() (articles []Article, err error) {
	log.Println("ListArticles() called")
	ctx := context.Background()
	client := getDBClient()

	iter := client.Collection("articles").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		var a Article
		doc.DataTo(&a)
		articles = append(articles, a)
	}

	return
}

// GetPostByID return a post matching the postID
func GetArticleByID(postID string) (post Post, err error) {

	return
}

// CreatePost inserts a new post to the database
func CreateArticle(newPost Post) (err error) {
	return
}
