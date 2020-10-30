package model

import (
	"html/template"
	"log"
	"time"
)

// A Post structure is for an article
type Post struct {
	ID              int
	Title           string
	Summary         string
	ContentMarkDown string
	ContentHTML     string
	Slug            string
	DateCreated     *time.Time

	// Not in database
	DisplayHTML template.HTML
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
		err = rows.Scan(&post.ID, &post.Title, &post.Summary, &post.ContentMarkDown, &post.ContentHTML, &post.Slug, &post.DateCreated)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return
}

// GetPostByID return an article matching the id
func GetPostBySlug(slug string) (post Post, err error) {
	log.Println("GetPostBySlug() called")
	client := getDBClient()
	defer client.Close()

	stmt, err := client.Prepare("select * from Post where Slug = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(slug).Scan(&post.ID, &post.Title, &post.Summary, &post.ContentMarkDown, &post.ContentHTML, &post.Slug, &post.DateCreated)
	if err != nil {
		log.Fatal(err)
	}

	post.DisplayHTML = template.HTML(post.ContentHTML)

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
