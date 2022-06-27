package images

import (
	"bytes"
	"image"
	"image/color"
	"io/ioutil"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pkg/errors"
)

func FromFile(path string) (*ebiten.Image, error) {
	b, err := ioutil.ReadFile(path)

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decode the image in path %q", path)
	}

	return ebiten.NewImageFromImage(img), nil
}

func CenterPosition(image *ebiten.Image, width, height int) (int, int) {
	iw, ih := image.Size()
	w := (width / 2) - (iw / 2)
	h := (height / 2) - (ih / 2)
	return w, h
}

func CircleImage(radius int, color color.Color) (*ebiten.Image, error) {
	c := radius * 2
	cf := float64(c)
	dc := gg.NewContext(c*2, c*2)
	dc.DrawCircle(cf, cf, cf)
	dc.SetColor(color)
	dc.Fill()

	return contextToImage(dc)
}

func RoundedRectImage(width int, height int, color color.Color) (*ebiten.Image, error) {
	dc := gg.NewContext(width, height)
	dc.DrawRoundedRectangle(0, 0, float64(width), float64(height), 20)
	dc.SetColor(color)
	dc.Fill()

	return contextToImage(dc)
}

func contextToImage(dc *gg.Context) (*ebiten.Image, error) {
	b := []byte{}
	buf := bytes.NewBuffer(b)
	err := dc.EncodePNG(buf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode image")
	}

	image, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode image")
	}

	return ebiten.NewImageFromImage(image), nil
}
