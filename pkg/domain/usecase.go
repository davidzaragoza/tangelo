package domain

import (
	"bytes"
	"fmt"
	"image"

	// Allow these image formats
	"image/draw"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"

	"log"
	"mime/multipart"
)

const (
	croppedBaseURL   = "http://localhost:8080/api/v1/cropped"
	verticalSplits   = 2
	horizontalSplits = 2
)

type UseCase struct {
	rep Repository
}

func NewUseCase(rep Repository) *UseCase {
	return &UseCase{rep: rep}
}

func (uc *UseCase) CropImage(name string, original multipart.File) ([]string, error) {
	log.Print("cropping the image")
	img, _, err := image.Decode(original)
	if err != nil {
		return nil, err
	}
	result := []string{}

	croppedWidth := img.Bounds().Dx() / verticalSplits
	croppedHeight := img.Bounds().Dy() / horizontalSplits

	for i := 0; i < horizontalSplits; i++ {
		for j := 0; j < verticalSplits; j++ {
			rectangle := image.Rectangle{image.Point{0, 0}, image.Point{croppedWidth, croppedHeight}}
			rgba := image.NewRGBA(rectangle)
			draw.Draw(rgba, rectangle.Bounds(), img, image.Point{j * croppedWidth, i * croppedHeight}, draw.Src)
			bytes, err := uc.getImageBytes(rgba)
			if err != nil {
				return nil, err
			}
			croppedName := fmt.Sprintf("%s/%s_%d_%d.jpg", croppedBaseURL, name, i, j)
			// uc.saveToDisk(fmt.Sprintf("%s_%d_%d.jpg", name, i, j), rgba)
			if err := uc.rep.SaveImage(croppedName, bytes); err != nil {
				return nil, err
			}
			result = append(result, croppedName)
		}
	}
	return result, nil
}

// func (uc *UseCase) saveToDisk(name string, rgba *image.RGBA) error {
// 	f, err := os.Create(name)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	var opt jpeg.Options
// 	opt.Quality = 100

// 	if err := jpeg.Encode(f, rgba, &opt); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (uc *UseCase) getImageBytes(rgba *image.RGBA) ([]byte, error) {
	var opt jpeg.Options
	opt.Quality = 100

	var result bytes.Buffer
	if err := jpeg.Encode(&result, rgba, &opt); err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}

func (uc *UseCase) GetImage(name string) ([]byte, error) {
	log.Printf("Obtaining image %s", name)
	croppedName := fmt.Sprintf("%s/%s", croppedBaseURL, name)
	return uc.rep.GetImage(croppedName)
}
