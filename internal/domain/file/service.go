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
	return nil, nil
}

func (s *service) GetFile(ctx context.Context, uuid string) *File {
	return s.storage.GetOne(uuid)
}
