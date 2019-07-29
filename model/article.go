package model

import (
	"log"

	"html/template"

	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

// Article Struct
type ArticleDB struct {
	ID       string `firestore:"id,omitempty"`
	Title    string `firestore:"title,omitempty"`
	Markdown string `firestore:"markdown,omitempty"`
	Content  string `firestore:"content,omitempty"`
	Slug     string `firestore:"slug,omitempty"`
	Created  string `firestore:"created,omitempty"`
}

type Article struct {
	ID       string
	Title    string
	Markdown string
	Content  template.HTML
	Slug     string
	Created  string
}

func convertArticleDBToArticle(adb ArticleDB) (a Article) {
	a.ID = adb.ID
	a.Title = adb.Title
	a.Markdown = adb.Markdown
	a.Content = template.HTML(adb.Content)
	a.Slug = adb.Slug
	a.Created = adb.Created
	return
}

// List all articles in the database
func ListArticles() (articles []Article, err error) {
	log.Println("ListArticles() called")
	ctx := context.Background()
	client := getDBClient()

	iter := client.Collection("articles").Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		var adb ArticleDB
		doc.DataTo(&adb)

		articles = append(articles, convertArticleDBToArticle(adb))
	}

	return
}

// GetArticleByID return an article matching the id
func GetArticleByID(id string) (at Article, err error) {
	log.Println("GetArticleById() called")
	ctx := context.Background()
	client := getDBClient()

	ats := client.Collection("articles")
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
		var adb ArticleDB
		doc.DataTo(&adb)
		at = convertArticleDBToArticle(adb)
	}
	return
}

// CreateArticle inserts a new article to the database
func CreateArticle(a ArticleDB) (err error) {
	log.Println("CreateArticle() called")
	ctx := context.Background()
	client := getDBClient()

	_, _, err = client.Collection("articles").Add(ctx, a)
	if err != nil {
		log.Fatalf("Failed to create article: %v", err)
	}

	return
}

// UpdateArticle update an article exists in the database
func UpdateArticle(a ArticleDB) (err error) {
	log.Println("UpdateArticle() called")
	ctx := context.Background()
	client := getDBClient()

	ats := client.Collection("articles")
	q := ats.Where("slug", "==", a.Slug)
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
			"title":   a.Title,
			"content": a.Content}, firestore.MergeAll)
	}

	if err != nil {
		log.Fatalf("Failed to update article: %v", err)
	}

	return
}
