package file

import (
	"github.com/kodaf1/go-cloud-storage/internal/domain/file"
	"go.mongodb.org/mongo-driver/mongo"
)

type fileStorage struct {
	db *mongo.Database
}

func NewStorage(db *mongo.Database) file.Storage {
	return &fileStorage{db: db}
}

func (bs *fileStorage) GetOne(uuid string) (*file.File, error) {
	return nil, nil
}
func (bs *fileStorage) GetAll(limit, offset int) []*file.File {
	return nil
}
func (bs *fileStorage) Create(book *file.File) *file.File {
	return nil
}
func (bs *fileStorage) Delete(book *file.File) error {
	return nil
}
