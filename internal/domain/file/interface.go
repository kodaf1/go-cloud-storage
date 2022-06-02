package file

import (
	"context"
	"mime/multipart"
)

type Service interface {
	UploadFile(ctx context.Context, dto *UploadFileDTO) (*File, error)
	GetFile(ctx context.Context, uuid string) *File
}

type Storage interface {
	GetOne(uuid string) *File
	GetAll(limit, offset int) []*File
	Create(user *File) *File
	Delete(user *File) error
}

type S3 interface {
	GetBucket() string
	PutObject(ctx context.Context, file *multipart.FileHeader, filename string) error
}
