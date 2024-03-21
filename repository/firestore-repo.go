package repository

import (
	"context"
	"go-mux-crash-course/entity"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type repo struct {
	CollectionName string
}

func NewFirestoreRepository(collectionName string) PostRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

const (
	projectID string = "go-mux-crash-course-519d4"
)

func (r *repo) Save(post *entity.Post) (ID int64, err error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return
	}

	_, _, err = client.Collection(r.CollectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	defer client.Close()
	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return
	}
	return
}

func (r *repo) FindAll() (posts []entity.Post, err error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return
	}

	defer client.Close()
	it := client.Collection(r.CollectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return
}

func (r *repo) FindByID(ID string) (post entity.Post, err error) {
	// NOT IMPLEMENTED
	return
}

func (r *repo) Delete(ID string) (err error) {
	return
}
