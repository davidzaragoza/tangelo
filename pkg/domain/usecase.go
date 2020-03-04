package domain

import (
	"fmt"
	"image"
	"os"

	// Allow these image formats
	"image/draw"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"

	"log"
	"mime/multipart"
)

const (
	verticalSplits   = 3
	horizontalSplits = 5
)

type UseCase struct {
}

func NewUseCase() *UseCase {
	return &UseCase{}
}

func (uc *UseCase) CropImage(original multipart.File) ([]string, error) {
	log.Print("cropping the image")
	img, _, err := image.Decode(original)
	if err != nil {
		return nil, err
	}

	croppedWidth := img.Bounds().Dx() / verticalSplits
	croppedHeight := img.Bounds().Dy() / horizontalSplits

	for i := 0; i < horizontalSplits; i++ {
		for j := 0; j < verticalSplits; j++ {
			rectangle := image.Rectangle{image.Point{0, 0}, image.Point{croppedWidth, croppedHeight}}
			rgba := image.NewRGBA(rectangle)
			draw.Draw(rgba, rectangle.Bounds(), img, image.Point{j * croppedWidth, i * croppedHeight}, draw.Src)
			if err := uc.saveImage(fmt.Sprintf("out_%d_%d.jpg", i, j), rgba); err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}

func (uc *UseCase) saveImage(name string, rgba *image.RGBA) error {
	out, err := os.Create(name)
	if err != nil {
		return err
	}
	defer out.Close()

	var opt jpeg.Options
	opt.Quality = 100

	jpeg.Encode(out, rgba, &opt)
	return nil
}
