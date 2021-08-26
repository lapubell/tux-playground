package game

import (
	"image"
	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type scene struct {
	Name              string
	Background        *ebiten.Image
	StartingPositionX uint
	StartingPositionY uint
	ExitingPositionX  uint
	ExitingPositionY  uint
	Exits             []sceneExit
}

type sceneExit struct {
	ToSceneIndex uint
	TopLeft      coordantant
	BottomRight  coordantant
}

type coordantant struct {
	x uint
	y uint
}

func initializeScenes() []scene {
	output := []scene{}
	// scene 1
	output = append(output, makeScene("../assets/scene-01.jpg", 540, 330, []sceneExit{
		{
			ToSceneIndex: 1,
			TopLeft:      coordantant{560, 318},
			BottomRight:  coordantant{640, 395},
		},
	}))

	// scene 2
	output = append(output, makeScene("../assets/scene-02.jpg", 40, 300, []sceneExit{
		{
			ToSceneIndex: 0,
			TopLeft:      coordantant{0, 284},
			BottomRight:  coordantant{80, 373},
		}, {
			ToSceneIndex: 2,
			TopLeft:      coordantant{417, 400},
			BottomRight:  coordantant{497, 480},
		},
	}))

	// scene 3
	output = append(output, makeScene("../assets/scene-03.jpg", 470, 20, []sceneExit{
		{
			ToSceneIndex: 1,
			TopLeft:      coordantant{468, 0},
			BottomRight:  coordantant{562, 80},
		},
	}))
	return output
}

func makeScene(imagePath string, startX uint, startY uint, exits []sceneExit) scene {
	sceneImage, _, _ := ebitenutil.NewImageFromFile(imagePath)
	s := scene{
		Name:              "Intro", // TODO: ditch this or make it a param
		StartingPositionX: startX,
		StartingPositionY: startY,
		Background:        sceneImage,
	}
	for _, e := range exits {
		s.Exits = append(s.Exits, e)
	}
	return s
}

func (g *Game) drawScene(screen *ebiten.Image) {
	screen.DrawImage(scenes[activeSceneIndex].Background.SubImage(image.Rect(0, 0, screenWidth, screenHeight)).(*ebiten.Image), (&ebiten.DrawImageOptions{}))
}

func checkSceneChange() {
	for _, exit := range scenes[activeSceneIndex].Exits {
		if tux.positionX >= float64(exit.TopLeft.x) &&
			tux.positionX+frameWidth <= float64(exit.BottomRight.x) &&
			tux.positionY >= float64(exit.TopLeft.y) &&
			tux.positionY+frameHeight <= float64(exit.BottomRight.y) {

			// set exiting positions
			// if these aren't set on the scene, then we can default to the
			// scene starting positions
			scenes[activeSceneIndex].ExitingPositionX = uint(tux.positionX)
			scenes[activeSceneIndex].ExitingPositionY = uint(tux.positionY)

			activeSceneIndex = exit.ToSceneIndex
			if scenes[activeSceneIndex].ExitingPositionX > 0 {
				offset := 0
				if tux.facingDirection == "L" {
					offset = -20
				}
				if tux.facingDirection == "R" {
					offset = 20
				}
				tux.positionX = float64(int(scenes[activeSceneIndex].ExitingPositionX) + offset)
			} else {
				tux.positionX = float64(scenes[activeSceneIndex].StartingPositionX)
			}

			if scenes[activeSceneIndex].ExitingPositionY > 0 {
				offset := 0
				if tux.facingDirection == "U" {
					offset = -20
				}
				if tux.facingDirection == "D" {
					offset = 20
				}
				tux.positionY = float64(int(scenes[activeSceneIndex].ExitingPositionY) + offset)
			} else {
				tux.positionY = float64(scenes[activeSceneIndex].StartingPositionY)
			}
			return
		}
	}
}
