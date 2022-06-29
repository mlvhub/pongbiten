package scenes

import (
	"github.com/mlvhub/pongbiten/pongbiten/colors"

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

	touchIDs []ebiten.TouchID
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

func (m *MenuScene) Update() (*SceneType, error) {
	nextSceneType := Game

	m.touchIDs = inpututil.AppendJustPressedTouchIDs(m.touchIDs[:0])
	touch := len(m.touchIDs) > 0

	if inpututil.KeyPressDuration(ebiten.KeySpace) > 0 || touch {
		return &nextSceneType, nil
	}
	return nil, nil
}

func (m *MenuScene) Draw(screen *ebiten.Image) {
	// title
	titleSize := text.BoundString(m.titleFont, m.title)
	titleX := (m.width / 2) - (titleSize.Dx() / 2)

	titleY := titleSize.Dy() * 2

	text.Draw(screen, m.title, m.titleFont, titleX, titleY, colors.Black)

	// controls
	player1 := "Player 1:\n \nW (up)\nS (down)\n"
	player2 := "Player 2:\n \nUp Arrow\nDown Arrow\n"
	player1Size := text.BoundString(m.font, player1)
	player2Size := text.BoundString(m.font, player2)

	player1X := (m.width / 2) - (player1Size.Dx() + (player1Size.Dx() / 2))
	player1Y := titleY + (titleY / 2) + (titleY / 4)
	text.Draw(screen, player1, m.font, player1X, player1Y, colors.Black)

	player2X := (m.width / 2) + (player2Size.Dx() / 4)
	text.Draw(screen, player2, m.font, player2X, player1Y, colors.Black)

	// start game
	msg := "Press 'Spacebar' \nor Tap to start the game!"
	msgSize := text.BoundString(m.font, msg)
	msgX := (m.width / 2) - (msgSize.Dx() / 2)

	msgY := player1Y + (player1Y / 2) + (player1Y / 4)

	text.Draw(screen, msg, m.font, msgX, msgY, colors.Black)
}
