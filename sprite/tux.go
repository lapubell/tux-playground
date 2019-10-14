package sprite

import (
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Tux - the main hero of our story
type Tux struct {
	currentX       float64
	currentY       float64
	currentChangeX float64
	currentChangeY float64
	isFlipped      bool
	artwork        string          // path to static asset
	Picture        pixel.Picture   // interface that does cool stuff
	State          string          // the "state" of our sprite - currently supports still|walking|jumping
	previousState  string          // the old "state" of our sprite
	frame          int             // the index of the current type of art to show
	speed          int64           // the number of milliseconds between frame changing
	currentScale   float64         // scale the artwork
	lastUpdated    time.Time       // timestamp to check if we need updating
	WalkingFrames  []*pixel.Sprite // all the mapping for the "walking" art
	JumpingFrames  []*pixel.Sprite // all the mapping for the "sliding" art
	currentFrame   *pixel.Sprite   // the sprite frame we want to draw
}

// Draw - draw our sprite
func (t *Tux) Draw(w *pixelgl.Window, m pixel.Matrix) {
	t.updateState()
	if t.isFlipped {
		m = m.ScaledXY(pixel.V(0, 0), pixel.V(-1, 1))
		m = m.Moved(pixel.V(math.Round(t.currentX+20), math.Round(t.currentY)))
	} else {
		m = m.Moved(pixel.V(math.Round(t.currentX), math.Round(t.currentY)))
	}
	t.currentFrame.Draw(w, m.Scaled(pixel.ZV, t.currentScale))
}

func (t *Tux) updateState() {
	if t.State != t.previousState {
		t.frame = 0
		t.previousState = t.State
	}

	if t.State == "still" {
		t.currentFrame = t.WalkingFrames[0]
	}

	if t.State == "walking" {
		if t.frame > len(t.WalkingFrames)-1 {
			t.frame = 0
		}
		t.currentFrame = t.WalkingFrames[t.frame]
	}

	if time.Now().Sub(t.lastUpdated).Milliseconds() > t.speed {
		t.frame++
		t.currentX += t.currentChangeX
		t.currentY += t.currentChangeY
		t.lastUpdated = time.Now()
	}
}

func (t *Tux) UpdateXAcceleration(x float64) {
	t.State = "walking"
	t.currentChangeX += x
	if t.currentChangeX > 3 {
		t.currentChangeX = 3
	}
	if t.currentChangeX < -3 {
		t.currentChangeX = -3
	}

	if t.currentChangeX > 0 {
		t.isFlipped = false
	}
	if t.currentChangeX < 0 {
		t.isFlipped = true
	}
}

func (t *Tux) UpdateYAcceleration(y float64) {
	t.State = "walking"
	t.currentChangeY += y
	if t.currentChangeY > 3 {
		t.currentChangeY = 3
	}
	if t.currentChangeY < -3 {
		t.currentChangeY = -3
	}
}

func (t *Tux) Decelerate() {
	if t.currentChangeX > 0 {
		t.currentChangeX -= 0.01
	} else if t.currentChangeX < 0 && t.currentChangeX+0.01 < 0 {
		t.currentChangeX += 0.01
	} else {
		t.currentChangeX = 0
	}

	if t.currentChangeY > 0 {
		t.currentChangeY -= 0.01
	} else if t.currentChangeY < 0 && t.currentChangeY+0.01 < 0 {
		t.currentChangeY += 0.01
	} else {
		t.currentChangeY = 0
	}

	if t.currentChangeX == 0 && t.currentChangeY == 0 {
		t.State = "still"
	}
}

// NewTux - return a new instance of the Tux sprite
func NewTux() Tux {
	t := Tux{}
	t.currentScale = 2.5
	t.speed = 50
	t.lastUpdated = time.Now()
	t.artwork = "../assets/Master System - Penguin Land - Overbite.png" // artwork found at https://www.spriters-resource.com/master_system/penguinland/sheet/85854/
	t.Picture, _ = loadPicture(t.artwork)
	// load up all the sprite sheet frames for the walking state
	t.WalkingFrames = append(t.WalkingFrames, pixel.NewSprite(t.Picture, pixel.R(56, 92, 74, 112)))
	t.WalkingFrames = append(t.WalkingFrames, pixel.NewSprite(t.Picture, pixel.R(74, 92, 91, 112)))
	t.WalkingFrames = append(t.WalkingFrames, pixel.NewSprite(t.Picture, pixel.R(91, 92, 108, 112)))
	t.WalkingFrames = append(t.WalkingFrames, pixel.NewSprite(t.Picture, pixel.R(108, 92, 125, 112)))
	t.WalkingFrames = append(t.WalkingFrames, pixel.NewSprite(t.Picture, pixel.R(125, 92, 143, 112)))
	t.WalkingFrames = append(t.WalkingFrames, pixel.NewSprite(t.Picture, pixel.R(143, 92, 160, 112)))
	t.WalkingFrames = append(t.WalkingFrames, pixel.NewSprite(t.Picture, pixel.R(160, 92, 176, 112)))

	t.State = "still"

	return t
}
