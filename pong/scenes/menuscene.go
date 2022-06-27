package scenes

import (
	"pong/pong/colors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type MenuScene struct {
	width  int
	height int

	title     string
	titleFont font.Face
	font      font.Face
}

func NewMenuScene(width int, height int, title string, titleFont font.Face, font font.Face) (*MenuScene, error) {
	return &MenuScene{
		width:  width,
		height: height,

		title:     title,
		titleFont: titleFont,
		font:      font,
	}, nil
}

func (g *MenuScene) Update() (*SceneType, error) {
	nextSceneType := Game
	if inpututil.KeyPressDuration(ebiten.KeySpace) > 0 {
		return &nextSceneType, nil
	}
	return nil, nil
}

func (g *MenuScene) Draw(screen *ebiten.Image) {
	// title
	titleSize := text.BoundString(g.titleFont, g.title)
	titleX := (g.width / 2) - (titleSize.Dx() / 2)

	titleY := titleSize.Dy() * 2

	text.Draw(screen, g.title, g.titleFont, titleX, titleY, colors.Black)

	// controls
	player1 := "Player 1:\n \nW (up)\nS (down)\n"
	player2 := "Player 2:\n \nUp Arrow\nDown Arrow\n"
	player1Size := text.BoundString(g.font, player1)
	player2Size := text.BoundString(g.font, player2)

	player1X := (g.width / 2) - (player1Size.Dx() + (player1Size.Dx() / 2))
	player1Y := titleY + (titleY / 2) + (titleY / 4)
	text.Draw(screen, player1, g.font, player1X, player1Y, colors.Black)

	player2X := (g.width / 2) + (player2Size.Dx() / 4)
	text.Draw(screen, player2, g.font, player2X, player1Y, colors.Black)

	// start game
	msg := "Press 'Spacebar' to start the game!"
	msgSize := text.BoundString(g.font, msg)
	msgX := (g.width / 2) - (msgSize.Dx() / 2)

	msgY := player1Y + (player1Y / 2) + (player1Y / 4)

	text.Draw(screen, msg, g.font, msgX, msgY, colors.Black)
}
