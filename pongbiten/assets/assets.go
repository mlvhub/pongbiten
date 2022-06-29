package assets

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"log"

	"github.com/mlvhub/pongbiten/pongbiten/colors"
	"github.com/mlvhub/pongbiten/pongbiten/images"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
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

//go:embed assets
var assetsRootFS embed.FS

type Assets struct {
	assetsFS     fs.FS
	AudioContext *audio.Context

	TitleFont font.Face
	Font      font.Face

	BallImage   *ebiten.Image
	PlayerImage *ebiten.Image

	HitAudioBytes []byte
}

func New() (*Assets, error) {
	assetsFS, err := fs.Sub(assetsRootFS, "assets")

	audioContext := audio.NewContext(48000)

	file, err := assetsFS.Open("hit.wav")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d, err := wav.Decode(audioContext, file)
	if err != nil {
		return nil, err
	}
	hitAudioBytes, err := ioutil.ReadAll(d)

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
		assetsFS:      assetsFS,
		AudioContext:  audioContext,
		TitleFont:     titleFont,
		Font:          font,
		BallImage:     ballImage,
		PlayerImage:   playerImage,
		HitAudioBytes: hitAudioBytes,
	}, nil

}

func (a *Assets) PlayHitSound() {
	player := audio.NewPlayerFromBytes(a.AudioContext, a.HitAudioBytes)
	player.Play()
}
