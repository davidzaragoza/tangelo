package domain

type StoredImage struct {
	ID      int
	Name    string
	Content []byte
}

type CropResult struct {
	ID           int
	OriginalFile StoredImage
	CroppedFiles []StoredImage
}

type DatabaseConfiguration struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
	SSLMode  string `json:"ssl"`
	Schema   string `json:"schema"`
}

type Configuration struct {
	Database DatabaseConfiguration `json:"database"`
}
