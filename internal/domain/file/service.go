package file

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type service struct {
	storage   Storage
	s3        S3
	URLPrefix string
}

func NewService(storage Storage, s3 S3, urlPrefix string) Service {
	return &service{storage, s3, urlPrefix}
}

func (s *service) UploadFile(ctx context.Context, dto *UploadFileDTO) (*File, error) {
	uid := uuid.NewString()

	err := s.s3.PutObject(ctx, dto.File, uid)
	if err != nil {
		return nil, err
	}

	file := &File{
		FileName: uid,
		URL:      fmt.Sprintf("%s%s", s.URLPrefix, uid),
		Size:     dto.File.Size,
	}

	return s.storage.Create(file)
}

func (s *service) GetFile(ctx context.Context, uuid string) (file *File, err error) {
	file, err = s.storage.GetOne(uuid)
	if err != nil {
		return file, err
	}

	file.URL = fmt.Sprintf("%s%s", s.URLPrefix, file.FileName)
	return file, err
}
