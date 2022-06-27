package ball

import (
	"math/rand"
	"pong/pong/player"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	image *ebiten.Image
	// position
	x int
	y int
	// velocity
	vx int
	vy int
	// direction
	dirX int
	dirY int

	// bounds
	maxX int
	maxY int

	// tracking collisions for speed
	collisions int

	// tracking the state to know whether or not the ball is out of bounds (a player lost)
	locationState LocationState
}

type LocationState int

const (
	WithinBounds LocationState = iota
	OutOfLeftBounds
	OutOfRightBounds
)

func New(maxX int, maxY int, image *ebiten.Image, initialX int, initialY int, initialVelocity int) *Ball {
	// random initial direction for the ball
	dirX := []int{-1, 1}[rand.Intn(2)]
	// always starts down
	dirY := 1

	return &Ball{
		// screen boundaries
		maxX: maxX,
		maxY: maxY,

		image: image,
		x:     initialX,
		y:     initialY,
		// which player gets the ball first
		vx: initialVelocity,
		// always down
		vy:            initialVelocity / 4,
		dirX:          dirX,
		dirY:          dirY,
		collisions:    0,
		locationState: WithinBounds,
	}
}

func (b *Ball) Update() {
	// TODO: should increase velocity the more times it hits a player
	b.x += (b.vx + b.collisions/2) * b.dirX
	b.y += (b.vy + b.collisions/2) * b.dirY

	w, h := b.Size()
	bottomY := b.y + h

	//check collision with top and bottom barriers
	if bottomY > b.maxY || b.y < 0 {
		b.dirY *= -1
		b.collisions++
	}

	// check location state
	if b.x-w < 0 {
		b.locationState = OutOfLeftBounds
	}
	if b.x+w > b.maxX {
		b.locationState = OutOfRightBounds
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.x), float64(b.y))

	screen.DrawImage(b.image, op)
}

func (b *Ball) HandleCollision(player *player.Player) bool {
	px, py := player.Position()
	pw, ph := player.Size()

	bw, bh := b.Size()

	// taken from https://developer.mozilla.org/en-US/docs/Games/Techniques/2D_collision_detection
	collided := b.x < px+pw &&
		b.x+bw > px &&
		b.y < py+ph &&
		bh+b.y > py

	if collided {
		b.dirX *= -1
	}

	return collided
}

func (b *Ball) LocationState() LocationState {
	return b.locationState
}

func (b *Ball) Size() (int, int) {
	return b.image.Size()
}
