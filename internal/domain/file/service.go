package file

import (
	"context"
	"github.com/google/uuid"
)

type service struct {
	storage Storage
	s3      S3
}

func NewService(storage Storage, s3 S3) Service {
	return &service{storage, s3}
}

func (s *service) UploadFile(ctx context.Context, dto *UploadFileDTO) (*File, error) {
	uid := uuid.NewString()

	err := s.s3.PutObject(ctx, dto.File, uid)
	if err != nil {
		return nil, err
	}

	file := &File{
		FileName: uid,
		Size:     dto.File.Size,
	}

	return s.storage.Create(file)
}

func (s *service) GetFile(ctx context.Context, uuid string) (*File, error) {
	return s.storage.GetOne(uuid)
}
