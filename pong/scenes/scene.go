package scenes

import "github.com/hajimehoshi/ebiten/v2"

type SceneType int

const (
	Menu SceneType = iota
	Game
	Results
)

type Scene interface {
	Update() (*SceneType, error)
	Draw(*ebiten.Image)
}
