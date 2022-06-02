package file

import (
	"context"
	"mime/multipart"
)

type Service interface {
	UploadFile(ctx context.Context, dto *UploadFileDTO) (*File, error)
	GetFile(ctx context.Context, uuid string) (*File, error)
}

type Storage interface {
	GetOne(uuid string) (*File, error)
	Create(file *File) (*File, error)
}

type S3 interface {
	GetBucket() string
	PutObject(ctx context.Context, file *multipart.FileHeader, filename string) error
}
