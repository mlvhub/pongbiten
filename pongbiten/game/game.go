package game

import (
	"github.com/mlvhub/pongbiten/pongbiten/assets"
	"github.com/mlvhub/pongbiten/pongbiten/colors"
	"github.com/mlvhub/pongbiten/pongbiten/scenes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pkg/errors"
)

type Game struct {
	width  int
	height int

	assets *assets.Assets

	currentSceneType scenes.SceneType

	menuScene    *scenes.MenuScene
	gameScene    *scenes.GameScene
	resultsScene *scenes.ResultsScene
}

func New(width int, height int, title string, assets *assets.Assets) (*Game, error) {
	initialSceneType := scenes.Menu

	menuScene, err := scenes.NewMenuScene(width, height, title, assets.TitleFont, assets.Font)
	if err != nil {
		return nil, err
	}

	return &Game{
		width:            width,
		height:           height,
		assets:           assets,
		currentSceneType: initialSceneType,
		menuScene:        menuScene,
		gameScene:        nil,
		resultsScene:     nil,
	}, nil
}

func (g *Game) Update() error {
	switch g.currentSceneType {
	case scenes.Menu:
		newSceneType, err := g.menuScene.Update()
		if err != nil {
			return errors.Wrap(err, "menu scene update")
		}

		if newSceneType != nil {
			gameScene, err := scenes.NewGameScene(g.width, g.height, g.assets)
			if err != nil {
				return errors.Wrap(err, "creating game scene")
			}
			g.menuScene = nil
			g.gameScene = gameScene
			g.resultsScene = nil
			g.currentSceneType = *newSceneType
		}
	case scenes.Game:
		newSceneType, err := g.gameScene.Update()
		if err != nil {
			return errors.Wrap(err, "game scene update")
		}

		if newSceneType != nil {
			resultsScene, err := scenes.NewResultsScene(g.width, g.height, g.assets.TitleFont, g.assets.Font, g.gameScene.Results())
			if err != nil {
				return errors.Wrap(err, "creating results scene")
			}
			g.menuScene = nil
			g.gameScene = nil
			g.resultsScene = resultsScene
			g.currentSceneType = *newSceneType
		}
	case scenes.Results:
		newSceneType, err := g.resultsScene.Update()
		if err != nil {
			return errors.Wrap(err, "results scene update")
		}

		if newSceneType != nil {
			gameScene, err := scenes.NewGameScene(g.width, g.height, g.assets)
			if err != nil {
				return errors.Wrap(err, "creating game scene")
			}
			g.menuScene = nil
			g.gameScene = gameScene
			g.resultsScene = nil
			g.currentSceneType = *newSceneType
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// background
	screen.Fill(colors.Grey)

	switch g.currentSceneType {
	case scenes.Menu:
		g.menuScene.Draw(screen)
	case scenes.Game:
		g.gameScene.Draw(screen)
	case scenes.Results:
		g.resultsScene.Draw(screen)

	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}
