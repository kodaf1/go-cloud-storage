package file

import (
	"context"
	"errors"
	"fmt"
	"github.com/kodaf1/go-cloud-storage/internal/domain/file"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type fileStorage struct {
	collection *mongo.Collection
}

func NewStorage(db *mongo.Database, collection string) file.Storage {
	return &fileStorage{collection: db.Collection(collection)}
}

func (fs *fileStorage) GetOne(uuid string) (f *file.File, err error) {
	objectID, err := primitive.ObjectIDFromHex(uuid)
	if err != nil {
		return f, fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": objectID}

	result := fs.collection.FindOne(context.Background(), filter)

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return f, fmt.Errorf("not found")
		}
		return f, fmt.Errorf("failed to execute query. error: %w", err)
	}

	if err = result.Decode(&f); err != nil {
		return f, fmt.Errorf("failed to decode document. error: %w", err)
	}

	return f, nil
}

func (fs *fileStorage) Create(file *file.File) (*file.File, error) {
	result, err := fs.collection.InsertOne(context.Background(), file)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		file.UUID = oid.Hex()
		return file, nil
	}
	return nil, fmt.Errorf("failed to convet objectid to hex")
}
