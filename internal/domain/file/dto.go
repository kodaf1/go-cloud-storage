package file

import "mime/multipart"

type UploadFileDTO struct {
	File *multipart.FileHeader
}
