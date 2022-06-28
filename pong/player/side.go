package player

import "github.com/hajimehoshi/ebiten/v2"

type side interface {
	ShouldMoveUp(x, y, width, height int) bool
	ShouldMoveDown(x, y, width, height int) bool
	UpKey() ebiten.Key
	DownKey() ebiten.Key
}

type left struct{}

func (l *left) ShouldMoveUp(x, y, width, height int) bool {
	// top left
	return x < width/2 && y < height/2
}
func (l *left) ShouldMoveDown(x, y, width, height int) bool {
	// bottom left
	return x < width/2 && y > height/2
}
func (l *left) UpKey() ebiten.Key {
	return ebiten.KeyW
}
func (l *left) DownKey() ebiten.Key {
	return ebiten.KeyS
}

func Left() side {
	return &left{}
}

type right struct{}

func (r *right) ShouldMoveUp(x, y, width, height int) bool {
	// top right
	return x > width/2 && y < height/2
}
func (r *right) ShouldMoveDown(x, y, width, height int) bool {
	// bottom right
	return x > width/2 && y > height/2
}
func (r *right) UpKey() ebiten.Key {
	return ebiten.KeyArrowUp
}
func (r *right) DownKey() ebiten.Key {
	return ebiten.KeyArrowDown
}

func Right() side {
	return &right{}
}
