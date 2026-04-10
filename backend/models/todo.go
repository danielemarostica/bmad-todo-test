package models

import (
	"context"
	"time"

	"bmad-todo-test/handlers"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoTodoStore struct {
	Collection *mongo.Collection
}

func NewTodoStore(collection *mongo.Collection) *MongoTodoStore {
	return &MongoTodoStore{Collection: collection}
}

func (s *MongoTodoStore) Create(text string) (handlers.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := bson.M{
		"text":      text,
		"completed": false,
		"createdAt": time.Now().UTC(),
	}

	result, err := s.Collection.InsertOne(ctx, doc)
	if err != nil {
		return handlers.Todo{}, err
	}

	id := result.InsertedID.(bson.ObjectID)
	return handlers.Todo{
		ID:        id.Hex(),
		Text:      text,
		Completed: false,
		CreatedAt: doc["createdAt"].(time.Time),
	}, nil
}
