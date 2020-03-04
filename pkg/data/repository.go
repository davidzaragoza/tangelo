package data

import "github.com/davidzaragoza/tangelo/pkg/domain"

type Repository struct {
}

func NewRepository() domain.Repository {
	return &Repository{}
}

func (rep *Repository) SaveImage(url string, content []byte) error {
	return nil
}

func (rep *Repository) GetImage(url string) ([]byte, error) {
	return nil, nil
}
