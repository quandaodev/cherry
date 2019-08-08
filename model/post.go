package model

import (
	"log"

	"html/template"

	"cloud.google.com/go/firestore"
	"github.com/quandaodev/cherry/utils"
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
	utils.LogInfo("ListPosts() called")
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
	utils.LogInfo("GetPostById() called")
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
			utils.LogError("Failed to iterate: ", err)
		}
		var pdb PostDB
		doc.DataTo(&pdb)
		at = convertPostDBToPost(pdb)
	}
	return
}

// CreatePost inserts a new post to the database
func CreatePost(p PostDB) (err error) {
	utils.LogInfo("CreatePost() called")
	ctx := context.Background()
	client := getDBClient()

	// TODO: check if slug is currently in database
	_, _, err = client.Collection("posts").Add(ctx, p)
	if err != nil {
		utils.LogError("Failed to create post: ", err)
	}

	return
}

// UpdatePost update an article exists in the database
func UpdatePost(p PostDB) (err error) {
	utils.LogInfo("UpdatePost() called")
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
			utils.LogError("Failed to iterate: ", err)
		}
		log.Println("Come here")
		_, err = doc.Ref.Set(ctx, map[string]interface{}{
			"title":   p.Title,
			"slug":    p.Slug,
			"content": p.Content,
			"html":    p.HTML}, firestore.MergeAll)
	}

	if err != nil {
		utils.LogError("Failed to update post: ", err)
	}

	return
}

// UpdatePost update an article exists in the database
func DeletePost(ID string) (err error) {
	utils.LogInfo("DeletePost() called")
	ctx := context.Background()
	client := getDBClient()

	pts := client.Collection("posts")
	q := pts.Where("slug", "==", ID)
	iter := q.Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			utils.LogError("Failed to iterate: ", err)
		}
		_, err = doc.Ref.Delete(ctx)
	}

	if err != nil {
		utils.LogError("Failed to delete post: ", err)
	}

	return
}
