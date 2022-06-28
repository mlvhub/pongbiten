package scenes

import (
	"fmt"
	"pong/pong/assets"
	"pong/pong/ball"
	"pong/pong/colors"
	"pong/pong/images"
	"pong/pong/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	ballInitialVelocity = 8

	player1Name = "Player 1"
	player2Name = "Player 2"

	playerOffset   = 20
	playerVelocity = 8

	maxPlayerScore = 5

	// ebiten ensures update is called 60 times a second, so a counter of 60 is one second
	maxReadyTimer = 60
)

type GameScene struct {
	width  int
	height int

	// convenience to reinit ball
	ballCenterX int
	ballCenterY int

	assets *assets.Assets

	ball *ball.Ball

	player1      *player.Player
	player1Score int
	player2      *player.Player
	player2Score int

	readyTimer        int
	readyTimerEnabled bool
}

type GameResults struct {
	player1Score int
	player1Name  string

	player2Score int
	player2Name  string
}

// TODO: shouldn't receive assets directly
func NewGameScene(width, height int, assets *assets.Assets) (*GameScene, error) {
	ballX, ballY := images.CenterPosition(assets.BallImage, width, height)
	ball := ball.New(width, height, assets.BallImage, ballX, ballY, ballInitialVelocity, assets)

	_, playerY := images.CenterPosition(assets.PlayerImage, width, height)

	player1 := player.New(player1Name, width, height, assets.PlayerImage, playerOffset, playerY, playerVelocity, player.Left())

	w, _ := assets.PlayerImage.Size()
	player2X := width - playerOffset - w
	player2 := player.New(player2Name, width, height, assets.PlayerImage, player2X, playerY, playerVelocity, player.Right())

	return &GameScene{
		width:        width,
		height:       height,
		ballCenterX:  ballX,
		ballCenterY:  ballY,
		assets:       assets,
		ball:         ball,
		player1:      player1,
		player1Score: 0,
		player2:      player2,
		player2Score: 0,
		// We'll start the timer straight away to give players a chance to get ready
		readyTimer:        maxReadyTimer,
		readyTimerEnabled: true,
	}, nil
}

func (g *GameScene) Update() (*SceneType, error) {
	// if we're showing the ready text we shouldn't update anything
	if g.readyTimer > 0 {
		g.readyTimer--
		return nil, nil
	}

	// if the timer was enabled and now it's zero, it means we need to spawn a new ball
	if g.readyTimer <= 0 && g.readyTimerEnabled {
		// reset the ball
		// reset the timer
		g.readyTimer = 0
		g.readyTimerEnabled = false
	}

	g.ball.Update()
	g.player1.Update()
	g.player2.Update()

	// check for collisions with the players
	g.ball.HandleCollision(g.player1)
	g.ball.HandleCollision(g.player2)

	scored := false
	// check if a player has scored
	switch g.ball.LocationState() {
	// player 1 scored
	case ball.OutOfRightBounds:
		scored = true
		g.player1Score++
	// player 2 scored
	case ball.OutOfLeftBounds:
		scored = true
		g.player2Score++
	}
	if scored {
		g.readyTimer = maxReadyTimer
		g.readyTimerEnabled = true
		// we reset the ball as early as possible to avoid weird stuff when we start drawing the ball again
		g.ball = ball.New(g.width, g.height, g.assets.BallImage, g.ballCenterX, g.ballCenterY, ballInitialVelocity, g.assets)
	}

	if g.player1Score >= maxPlayerScore || g.player2Score >= maxPlayerScore {
		newScene := Results
		return &newScene, nil
	}

	return nil, nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	if g.readyTimer == 0 {
		g.ball.Draw(screen)
	} else {
		// start game
		msg := "Get ready!"
		msgSize := text.BoundString(g.assets.Font, msg)
		msgX := (g.width / 2) - (msgSize.Dx() / 2)
		msgY := (g.height / 2) - (msgSize.Dy() / 2)
		text.Draw(screen, msg, g.assets.Font, msgX, msgY, colors.Black)
	}

	g.player1.Draw(screen)
	g.player2.Draw(screen)

	offset := 20
	p1 := fmt.Sprintf("%s: %d", player1Name, g.player1Score)
	p1Size := text.BoundString(g.assets.Font, p1)
	p1X := offset
	p1Y := p1Size.Dy() + offset
	text.Draw(screen, p1, g.assets.Font, p1X, p1Y, colors.Black)

	p2 := fmt.Sprintf("%s: %d", player2Name, g.player2Score)
	p2Size := text.BoundString(g.assets.Font, p2)
	p2X := g.width - p2Size.Dx() - offset
	p2Y := p2Size.Dy() + offset
	text.Draw(screen, p2, g.assets.Font, p2X, p2Y, colors.Black)
}

func (g *GameScene) Results() *GameResults {
	return &GameResults{
		player1Score: g.player1Score,
		player1Name:  player1Name,

		player2Score: g.player2Score,
		player2Name:  player2Name,
	}
}
