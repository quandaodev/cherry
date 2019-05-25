package model

import (
	"log"

	"golang.org/x/net/context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func getDBClient() (client *firestore.Client) {

	ctx := context.Background()
	opt := option.WithCredentialsFile("serviceAccountKey.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}

	client, err = app.Firestore(ctx)

	return client
}
