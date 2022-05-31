package file

import "context"


type Service interface {
	UploadFile(ctx context.Context, dto *UploadFileDTO) *File
	GetFile(ctx context.Context, uuid string) (*File, error)
}

type Storage interface {
	GetOne(uuid string) (*File, error)
	GetAll(limit, offset int) []*File
	Create(user *File) *File
	Delete(user *File) error
}
