package model

import (
	"log"

	"html/template"

	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

// Article Struct
type PostDB struct {
	ID      string `firestore:"id,omitempty"`
	Title   string `firestore:"title,omitempty"`
	Content string `firestore:"content,omitempty"`
	HTML    string `firestore:"html,omitempty"`
	Slug    string `firestore:"slug,omitempty"`
	Created string `firestore:"created,omitempty"`
}

type Post struct {
	ID      string
	Title   string
	Content string
	HTML    template.HTML
	Slug    string
	Created string
}

func convertPostDBToPost(pdb PostDB) (p Post) {
	p.ID = pdb.ID
	p.Title = pdb.Title
	p.Content = pdb.Content
	p.HTML = template.HTML(pdb.HTML)
	p.Slug = pdb.Slug
	p.Created = pdb.Created
	return
}

// List all posts in the database
func ListPosts() (posts []Post, err error) {
	log.Println("ListPosts() called")
	ctx := context.Background()
	client := getDBClient()

	iter := client.Collection("posts").Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		var pdb PostDB
		doc.DataTo(&pdb)

		posts = append(posts, convertPostDBToPost(pdb))
	}

	return
}

// GetPostByID return an article matching the id
func GetPostByID(id string) (at Post, err error) {
	log.Println("GetPostById() called")
	ctx := context.Background()
	client := getDBClient()

	ats := client.Collection("posts")
	q := ats.Where("slug", "==", id)
	iter := q.Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		var pdb PostDB
		doc.DataTo(&pdb)
		at = convertPostDBToPost(pdb)
	}
	return
}

// CreatePost inserts a new post to the database
func CreatePost(p PostDB) (err error) {
	log.Println("CreatePost() called")
	ctx := context.Background()
	client := getDBClient()

	// TODO: check if slug is currently in database
	_, _, err = client.Collection("posts").Add(ctx, p)
	if err != nil {
		log.Fatalf("Failed to create post: %v", err)
	}

	return
}

// UpdatePost update an article exists in the database
func UpdatePost(p PostDB) (err error) {
	log.Println("UpdatePost() called")
	ctx := context.Background()
	client := getDBClient()

	// TODO: check if slug currently in database
	pts := client.Collection("posts")
	q := pts.Where("slug", "==", p.Slug)
	iter := q.Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		log.Println("Come here")
		_, err = doc.Ref.Set(ctx, map[string]interface{}{
			"title":   p.Title,
			"slug":    p.Slug,
			"content": p.Content,
			"html":    p.HTML}, firestore.MergeAll)
	}

	if err != nil {
		log.Fatalf("Failed to update post: %v", err)
	}

	return
}
