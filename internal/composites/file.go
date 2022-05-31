package composites

import (
	"github.com/kodaf1/go-cloud-storage/internal/adapters/api"
	file3 "github.com/kodaf1/go-cloud-storage/internal/adapters/api/file"
	file2 "github.com/kodaf1/go-cloud-storage/internal/adapters/db/file"
	"github.com/kodaf1/go-cloud-storage/internal/domain/file"
)

type FileComposite struct {
	Storage file.Storage
	Service file.Service
	Handler api.Handler
}

func NewFileComposite(mongoComposite *MongoDBComposite) (*FileComposite, error) {
	storage := file2.NewStorage(mongoComposite.db)
	service := file.NewService(storage)
	handler := file3.NewHandler(service)
	return &FileComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
