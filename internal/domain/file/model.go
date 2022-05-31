package file

type File struct {
	UUID     string `json:"uuid"`
	FileName string `json:"filename"`
	MimeType string `json:"mime_type"`
	Size     int    `json:"size"`
}
