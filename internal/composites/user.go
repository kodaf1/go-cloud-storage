package composites

import (
	"github.com/kodaf1/go-cloud-storage/internal/adapters/api"
	"github.com/kodaf1/go-cloud-storage/internal/domain/user"
)

type UserComposite struct {
	Storage user.Storage
	Service user.Service
	Handler api.Handler
}

func NewUserComposite(mongoComposite *MongoDBComposite) (*UserComposite, error) {
	storage := user.NewStorage(mongoComposite.db)
	service := user.NewService(storage)
	handler := user.NewHandler(service)
	return &UserComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
