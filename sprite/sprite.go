package sprite

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Sprite a thing that can be drawn
type Sprite interface {
	Draw(*pixelgl.Window, pixel.Matrix)
}
