package model

import (
	"fmt"
	"log"
	"time"
)

// A Post structure is for an article
type Post struct {
	ID              int
	Title           string
	ContentMarkDown string
	ContentHTML     string
	Slug            string
	DateCreated     *time.Time
}

// List all posts in the database
func ListPosts() (posts []Post, err error) {
	log.Println("ListPosts() called")
	client := getDBClient()
	defer client.Close()

	rows, err := client.Query("select * from Post")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.Title, &post.ContentMarkDown, &post.ContentHTML, &post.Slug, &post.DateCreated)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(post.ID, post.Title, post.ContentHTML, post.DateCreated)
		posts = append(posts, post)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return
}

// GetPostByID return an article matching the id
func GetPostByID(id string) (at Post, err error) {
	log.Println("GetPostById() called")

	return
}

// CreatePost inserts a new post to the database
func CreatePost(p Post) (err error) {
	log.Println("CreatePost() called")

	return
}

// UpdatePost update an article exists in the database
func UpdatePost(p Post) (err error) {
	log.Println("UpdatePost() called")

	if err != nil {
		log.Fatalf("Failed to update post: %v", err)
	}

	return
}
