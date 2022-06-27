package assets

import (
	"log"
	"pong/pong/colors"
	"pong/pong/images"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	dpi           = 72
	titleFontSize = 80
	fontSize      = 30

	ballRadius = 8

	PlayerWidth  = 40
	PlayerHeight = 120
)

type Assets struct {
	TitleFont   font.Face
	Font        font.Face
	BallImage   *ebiten.Image
	PlayerImage *ebiten.Image
}

func New() (*Assets, error) {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}

	titleFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    titleFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	font, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	ballImage, err := images.CircleImage(ballRadius, colors.White)
	if err != nil {
		return nil, err
	}

	playerImage, err := images.RoundedRectImage(PlayerWidth, PlayerHeight, colors.Yellow)
	if err != nil {
		return nil, err
	}

	return &Assets{
		TitleFont:   titleFont,
		Font:        font,
		BallImage:   ballImage,
		PlayerImage: playerImage,
	}, nil

}
