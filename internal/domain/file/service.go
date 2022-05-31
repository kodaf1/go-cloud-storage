package file

import "context"

type service struct {
	storage       Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) UploadFile(ctx context.Context, dto *UploadFileDTO) *File {
	return nil
}

func (s *service) GetFile(ctx context.Context, uuid string) (*File, error) {
	return s.storage.GetOne(uuid)
}