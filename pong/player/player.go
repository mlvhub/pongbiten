package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	maxY int

	image *ebiten.Image

	x int
	y int

	vx int
	vy int

	upKey   ebiten.Key
	downKey ebiten.Key
}

func New(maxY int, image *ebiten.Image, initialX int, initialY int, velocity int, upKey ebiten.Key, downKey ebiten.Key) *Player {
	return &Player{
		maxY:    maxY,
		image:   image,
		x:       initialX,
		y:       initialY,
		vy:      velocity,
		upKey:   upKey,
		downKey: downKey,
	}
}

func (p *Player) Update() {
	if inpututil.KeyPressDuration(p.upKey) > 0 {
		p.y -= p.vy
	}
	if inpututil.KeyPressDuration(p.downKey) > 0 {
		p.y += p.vy
	}

	_, h := p.Size()
	bottomY := p.y + h

	if bottomY > p.maxY {
		p.y = p.maxY - h
	}

	if p.y < 0 {
		p.y = 0
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.x), float64(p.y))

	screen.DrawImage(p.image, op)
}

func (p *Player) Size() (int, int) {
	return p.image.Size()
}

func (p *Player) Position() (int, int) {
	return p.x, p.y
}
