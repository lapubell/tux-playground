package game

import "github.com/hajimehoshi/ebiten/v2"

type sprite struct {
	image           *ebiten.Image
	frameNumber     int
	facingDirection string
	isWalking       bool
	positionX       float64
	positionY       float64
}
