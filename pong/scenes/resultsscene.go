package scenes

import (
	"fmt"
	"pong/pong/colors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

const (
	title = ""
)

type ResultsScene struct {
	width  int
	height int

	titleFont font.Face
	font      font.Face

	results *GameResults

	touchIDs []ebiten.TouchID
}

func NewResultsScene(width int, height int, titleFont font.Face, font font.Face, results *GameResults) (*ResultsScene, error) {
	return &ResultsScene{
		width:  width,
		height: height,

		titleFont: titleFont,
		font:      font,

		results: results,
	}, nil
}

func (r *ResultsScene) Update() (*SceneType, error) {
	r.touchIDs = inpututil.AppendJustPressedTouchIDs(r.touchIDs[:0])
	touch := len(r.touchIDs) > 0

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || touch {
		newScene := Game
		return &newScene, nil
	}
	return nil, nil
}

func (r *ResultsScene) Draw(screen *ebiten.Image) {
	var title string
	if r.results.player1Score > r.results.player2Score {
		title = fmt.Sprintf("%s won!", r.results.player1Name)
	} else {
		title = fmt.Sprintf("%s won!", r.results.player2Name)
	}

	// title
	titleSize := text.BoundString(r.titleFont, title)
	titleX := (r.width / 2) - (titleSize.Dx() / 2)

	titleY := titleSize.Dy() * 2

	text.Draw(screen, title, r.titleFont, titleX, titleY, colors.Black)

	// start game
	msg := "Press 'Spacebar' or Tap to play again!"
	msgSize := text.BoundString(r.font, msg)
	msgX := (r.width / 2) - (msgSize.Dx() / 2)

	msgY := titleY + (titleY / 2) + (titleY / 4)

	text.Draw(screen, msg, r.font, msgX, msgY, colors.Black)
}
