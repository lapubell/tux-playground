package game

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func setupSprites() {
	tux.image, _, _ = ebitenutil.NewImageFromFile("../assets/tux.png")
	tux.facingDirection = "D"
	tux.positionX = 80
	tux.positionY = 300
	tuxDark.image, _, _ = ebitenutil.NewImageFromFile("../assets/tux-dark.png")
	tuxBlue.image, _, _ = ebitenutil.NewImageFromFile("../assets/tux-blue.png")
	tuxGreen.image, _, _ = ebitenutil.NewImageFromFile("../assets/tux-green.png")
	tuxRed.image, _, _ = ebitenutil.NewImageFromFile("../assets/tux-red.png")
}
