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
