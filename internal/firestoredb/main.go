package firestoredb

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// CreateClient creats a client to interact with the firestore
func CreateClient(ctx context.Context, projectID string) *firestore.Client {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Unable to create context: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

// Add creates a todo in the firestore collection
func Add(ctx context.Context, c *firestore.Client) {
	_, _, err := c.Collection("todos").Add(ctx, map[string]interface{}{
		"id":          1,
		"description": "clean the house",
	})
	if err != nil {
		log.Printf("Failed adding: %v", err)
		// l.Fatalf("Failed adding alovelace: %v", err)
	}
}
