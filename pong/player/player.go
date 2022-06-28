package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	Name string

	maxX int
	maxY int

	image *ebiten.Image

	x int
	y int

	vx int
	vy int

	side side

	// TODO: do we need to we track of this on every object or just scene?
	touchIDs []ebiten.TouchID
}

func New(name string, maxX int, maxY int, image *ebiten.Image, initialX int, initialY int, velocity int, side side) *Player {
	return &Player{
		Name: name,

		maxX:     maxX,
		maxY:     maxY,
		image:    image,
		x:        initialX,
		y:        initialY,
		vy:       velocity,
		side:     side,
		touchIDs: []ebiten.TouchID{},
	}
}

func (p *Player) Update() {
	p.touchIDs = inpututil.AppendJustPressedTouchIDs(p.touchIDs)

	relevantTouchIDs := []ebiten.TouchID{}
	for _, t := range p.touchIDs {
		if inpututil.TouchPressDuration(t) > 0 {
			relevantTouchIDs = append(relevantTouchIDs, t)
			x, y := ebiten.TouchPosition(t)

			// top left
			if p.side.ShouldMoveUp(x, y, p.maxX, p.maxY) {
				p.y -= p.vy
			}
			if p.side.ShouldMoveDown(x, y, p.maxX, p.maxY) {
				p.y += p.vy
			}
		}
	}
	p.touchIDs = relevantTouchIDs

	if inpututil.KeyPressDuration(p.side.UpKey()) > 0 {
		p.y -= p.vy
	}
	if inpututil.KeyPressDuration(p.side.DownKey()) > 0 {
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
