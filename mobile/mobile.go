package mobile

import (
	"log"
	"math/rand"
	"time"

	"github.com/mlvhub/pongbiten/pongbiten/assets"
	"github.com/mlvhub/pongbiten/pongbiten/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/mobile"
)

const (
	title  = "Pongbiten"
	width  = 1200
	height = 600
)

func init() {
	// pseudo-randomness
	rand.Seed(time.Now().UnixNano())

	// window config
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowSize(width, height)

	// create and load assets
	assets, err := assets.New()
	if err != nil {
		log.Fatalf("failed to create and load assets: %v", err)
	}

	// create a new game
	game, err := game.New(width, height, title, assets)
	if err != nil {
		log.Fatalf("failed creating a new game: %v", err)
	}

	// yourgame.Game must implement ebiten.Game interface.
	// For more details, see
	// * https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Game
	mobile.SetGame(game)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
