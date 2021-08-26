package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	count int
	keys  []ebiten.Key
}

func (g *Game) Update() error {
	g.count++
	g.keys = inpututil.PressedKeys()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawScene(screen)
	g.handleInput()
	g.updateSprites()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(tux.positionX, tux.positionY)

	sx, sy := 0, 0
	sx = frameOX + tux.frameNumber*frameWidth
	sy = frameOY
	if tux.frameNumber > 6 {
		offset := tux.frameNumber - 7
		sx = frameOX + offset*frameWidth
		sy = frameOY + frameHeight
	}

	screen.DrawImage(tux.image.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)

	g.checkCollisions()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Run() error {
	scenes = initializeScenes()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(windowTitle)
	setupSprites()

	ebiten.RunGame(g)

	return nil
}
