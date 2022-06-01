package composites

import (
	"github.com/kodaf1/go-cloud-storage/internal/adapters/api"
	file3 "github.com/kodaf1/go-cloud-storage/internal/adapters/api/file"
	file2 "github.com/kodaf1/go-cloud-storage/internal/adapters/db/file"
	file4 "github.com/kodaf1/go-cloud-storage/internal/adapters/s3/file"
	"github.com/kodaf1/go-cloud-storage/internal/domain/file"
)

type FileComposite struct {
	Storage file.Storage
	Service file.Service
	Handler api.Handler
	S3      file.S3
}

func NewFileComposite(mongoComposite *MongoDBComposite, s3Composite *S3Composite) (*FileComposite, error) {
	storage := file2.NewStorage(mongoComposite.db)
	service := file.NewService(storage)
	handler := file3.NewHandler(service)
	s3 := file4.NewS3(s3Composite.instance)
	return &FileComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
		S3:      s3,
	}, nil
}
