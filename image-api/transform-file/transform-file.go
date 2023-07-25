package transformfile

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"io"
)

func TransformFile(file io.Reader) ([]byte, error) {
	img, err := png.Decode(file)

	if err != nil {
		panic(err)
	}

	imgSize := img.Bounds().Size()

	rect := image.Rect(0, 0, imgSize.X, imgSize.Y)

	newImg := image.NewRGBA(rect)

	for x := 0; x <= newImg.Bounds().Max.X; x++ {

		for y := 0; y <= newImg.Bounds().Max.Y; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			newR := mirrorInt(int(r / 257))
			newG := mirrorInt(int(g / 257))
			newB := mirrorInt(int(b / 257))
			newColor := color.RGBA{byte(newR), byte(newG), byte(newB), byte(a)}
			newImg.Set(x, y, newColor)
		}
	}

	var buf bytes.Buffer

	if err = png.Encode(&buf, newImg); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func mirrorInt(n int) int {

	currentPercent := float32(n) / float32(255)
	mirrorPercent := 1 - currentPercent
	return int(255 * mirrorPercent)
}
