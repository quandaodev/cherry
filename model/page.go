package model

import (
	"log"

	"html/template"

	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

// PageDB is a database struct for Page
type PageDB struct {
	ID      string `firestore:"id,omitempty"`
	Content string `firestore:"content,omitempty"`
	HTML    string `firestore:"html,omitempty"`
}

// Page is a display struct for Page
type Page struct {
	ID      string
	Content string
	HTML    template.HTML
}

func convertPageDBToPage(pdb PageDB) (p Page) {
	p.ID = pdb.ID
	p.HTML = template.HTML(pdb.HTML)
	p.Content = pdb.Content
	return
}

// ListPages list all pages in the database
func ListPages() (pages []Page, err error) {
	log.Println("ListPages() called")
	ctx := context.Background()
	client := getDBClient()

	iter := client.Collection("pages").Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		var pdb PageDB
		doc.DataTo(&pdb)

		pages = append(pages, convertPageDBToPage(pdb))
	}

	return
}

// GetPageByID return an article matching the id
func GetPageByID(id string) (p Page, err error) {
	log.Println("GetPageById() called")
	ctx := context.Background()
	client := getDBClient()

	dsnap, _ := client.Collection("pages").Doc(id).Get(ctx)

	var pdb PageDB
	dsnap.DataTo(&pdb)
	p = convertPageDBToPage(pdb)

	return
}

// CreatePage inserts a new page to the database
func CreatePage(p PageDB) (err error) {
	log.Println("CreatePage() called")
	ctx := context.Background()
	client := getDBClient()

	_, _, err = client.Collection("pages").Add(ctx, p)
	if err != nil {
		log.Fatalf("Failed to create page: %v", err)
	}

	return
}

// UpdatePage update a page exists in the database
func UpdatePage(p PageDB) (err error) {
	log.Println("UpdatePage() called")
	ctx := context.Background()
	client := getDBClient()

	pts := client.Collection("pages")
	q := pts.Where("id", "==", p.ID)
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
		_, err = doc.Ref.Set(ctx, map[string]interface{}{
			"content": p.Content,
			"html":    p.HTML}, firestore.MergeAll)
	}

	if err != nil {
		log.Fatalf("Failed to update page: %v", err)
	}

	return
}
