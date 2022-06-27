package main

import (
	"log"
	"math/rand"
	"pong/pong/assets"
	"pong/pong/game"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	title  = "Pongbiten"
	width  = 1200
	height = 600
)

func main() {
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

	// run the game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("failed running the game")
	}
}
