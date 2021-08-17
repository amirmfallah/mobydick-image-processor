package lib

import (
	"image"
	"math"

	"github.com/nfnt/resize"
)

type point struct {
	x int
	y int
}

func SquareFitCrop(img image.Image) image.Image {
	imageW := float64(img.Bounds().Max.X)
	imageH := float64(img.Bounds().Max.Y)

	squareW := math.Min(imageH, imageW)
	diff := point{
		x: int(imageW - squareW),
		y: int(imageH - squareW),
	}

	crop0 := point{
		x: int(math.Floor(float64(diff.x) / 2)),
		y: int(math.Floor(float64(diff.y) / 2)),
	}

	crop1 := point{
		x: int(imageW) - int(math.Floor(float64(diff.x)/2)),
		y: int(imageH) - int(math.Floor(float64(diff.y)/2)),
	}

	cropped := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(crop0.x, crop0.y, crop1.x, crop1.y))
	return cropped
}

func ImageResize(img image.Image, width uint, height uint) image.Image {
	m := resize.Resize(width, height, img, resize.NearestNeighbor)
	return m
}
