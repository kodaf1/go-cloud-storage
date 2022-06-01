package file

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kodaf1/go-cloud-storage/internal/domain/file"
	"io"
)

type fileS3 struct {
	s3 *s3.S3
}

func (f fileS3) PutObject(file io.Reader) error {
	//TODO implement me
	panic("implement me")
}

func NewS3(s3 *s3.S3) file.S3 {
	return fileS3{s3: s3}
}
