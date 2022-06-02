package file

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kodaf1/go-cloud-storage/internal/domain/file"
	"mime"
	"mime/multipart"
	"strings"
)

type fileS3 struct {
	s3     *s3.S3
	bucket string
}

func (f fileS3) GetBucket() string {
	return f.bucket
}

func (f fileS3) PutObject(ctx context.Context, file *multipart.FileHeader, filename string) error {
	sFilename := strings.Split(file.Filename, ".")
	fileReader, err := file.Open()
	if err != nil {
		return err
	}

	_, err = f.s3.PutObjectWithContext(ctx, &s3.PutObjectInput{
		ACL:         aws.String("public-read"),
		Body:        fileReader,
		Bucket:      aws.String(f.bucket),
		Key:         aws.String(filename),
		ContentType: aws.String(mime.TypeByExtension("." + sFilename[len(sFilename)-1])),
	})

	return err
}

func NewS3(s3 *s3.S3, bucket string) file.S3 {
	return fileS3{s3, bucket}
}
