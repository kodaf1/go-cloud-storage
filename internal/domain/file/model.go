package file

type File struct {
	UUID     string `json:"uuid" bson:"_id,omitempty"`
	FileName string `json:"filename" bson:"file_name"`
	URL      string `json:"url,omitempty"`
	Size     int64  `json:"size" bson:"size"`
}
