package models

import (
	"context"
	"fmt"
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

type todoDoc struct {
	ID        bson.ObjectID `bson:"_id"`
	Text      string        `bson:"text"`
	Completed bool          `bson:"completed"`
	CreatedAt time.Time     `bson:"createdAt"`
}

func (s *MongoTodoStore) List() ([]handlers.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := s.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var docs []todoDoc
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	todos := make([]handlers.Todo, len(docs))
	for i, doc := range docs {
		todos[i] = handlers.Todo{
			ID:        doc.ID.Hex(),
			Text:      doc.Text,
			Completed: doc.Completed,
			CreatedAt: doc.CreatedAt,
		}
	}
	return todos, nil
}

func (s *MongoTodoStore) Update(id string, completed bool) (handlers.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return handlers.Todo{}, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"completed": completed}}

	result := s.Collection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		return handlers.Todo{}, result.Err()
	}

	var doc todoDoc
	if err := result.Decode(&doc); err != nil {
		return handlers.Todo{}, err
	}

	return handlers.Todo{
		ID:        doc.ID.Hex(),
		Text:      doc.Text,
		Completed: completed,
		CreatedAt: doc.CreatedAt,
	}, nil
}

func (s *MongoTodoStore) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id")
	}

	result, err := s.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("not found")
	}

	return nil
}
