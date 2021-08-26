package game

const (
	screenWidth  = 640
	screenHeight = 480

	frameOX     = 0
	frameOY     = 0
	frameWidth  = 64
	frameHeight = 64
	frameNum    = 3
	windowTitle = "Tux's Playground"

	tuxSpeed = 3
)

var (
	tux      sprite
	tuxDark  sprite
	tuxBlue  sprite
	tuxGreen sprite
	tuxRed   sprite

	// key codes that are used to toggle walking
	walkingKeys = []string{"ArrowUp", "ArrowDown", "ArrowRight", "ArrowLeft"}

	// scene stuff
	scenes           = []scene{} // check out initializeScenes()
	activeSceneIndex = 0
)
