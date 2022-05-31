package composites

import (
	"context"

	mongodb "github.com/kodaf1/go-cloud-storage/pkg/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBComposite struct {
	db *mongo.Database
}

func NewMongoDBComposite(ctx context.Context, Host, Port, Username, Password, Database, AuthSource string) (*MongoDBComposite, error) {
	client, err := mongodb.NewClient(ctx, Host, Port, Username, Password, Database, AuthSource)
	if err != nil {
		return nil, err
	}
	return &MongoDBComposite{db: client}, nil
}
