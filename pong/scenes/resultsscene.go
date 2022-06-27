package scenes

import (
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

func (g *ResultsScene) Update() (*SceneType, error) {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		newScene := Game
		return &newScene, nil
	}
	return nil, nil
}

func (g *ResultsScene) Draw(screen *ebiten.Image) {
	var title string
	if g.results.player1Score > g.results.player2Score {
		title = "Player 1 won!"
	} else {
		title = "Player 2 won!"
	}

	// title
	titleSize := text.BoundString(g.titleFont, title)
	titleX := (g.width / 2) - (titleSize.Dx() / 2)

	titleY := titleSize.Dy() * 2

	text.Draw(screen, title, g.titleFont, titleX, titleY, colors.Black)

	// start game
	msg := "Press 'Spacebar' to play again!"
	msgSize := text.BoundString(g.font, msg)
	msgX := (g.width / 2) - (msgSize.Dx() / 2)

	msgY := titleY + (titleY / 2) + (titleY / 4)

	text.Draw(screen, msg, g.font, msgX, msgY, colors.Black)
}
