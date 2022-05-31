package file

import (
	"io"
)

type UploadFileDTO struct {
	Name   string
	Size   int64
	Reader io.Reader
}
