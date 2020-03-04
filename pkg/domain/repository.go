package domain

type Repository interface {
	SaveImage(url string, content []byte) error
	GetImage(url string) ([]byte, error)
}
