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
	StartingPositionX int
	StartingPositionY int
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
	scene1Image, _, _ := ebitenutil.NewImageFromFile("../assets/scene-01.jpg")
	scene1Exit := sceneExit{
		ToSceneIndex: 1,
		TopLeft:      coordantant{622, 320},
		BottomRight:  coordantant{640, 395},
	}
	scene1 := scene{
		Name:              "Intro",
		StartingPositionX: 540,
		StartingPositionY: 330,
		Background:        scene1Image,
	}
	scene1.Exits = append(scene1.Exits, scene1Exit)
	scene2Image, _, _ := ebitenutil.NewImageFromFile("../assets/scene-02.jpg")
	scene2 := scene{
		Name:              "Scene 2",
		StartingPositionX: 40,
		StartingPositionY: 300,
		Background:        scene2Image,
	}
	scene2Exit := sceneExit{
		ToSceneIndex: 0,
		TopLeft:      coordantant{0, 320},
		BottomRight:  coordantant{30, 395},
	}
	scene2.Exits = append(scene2.Exits, scene2Exit)

	output = append(output, scene1)
	output = append(output, scene2)
	return output
}

func (g *Game) drawScene(screen *ebiten.Image) {
	screen.DrawImage(scenes[activeSceneIndex].Background.SubImage(image.Rect(0, 0, screenWidth, screenHeight)).(*ebiten.Image), (&ebiten.DrawImageOptions{}))
}

func checkSceneChange() {
	for _, exit := range scenes[activeSceneIndex].Exits {
		if tux.positionX+frameWidth > float64(exit.TopLeft.x) && tux.positionX < float64(exit.BottomRight.x) {
			activeSceneIndex = int(exit.ToSceneIndex)
			tux.positionX = float64(scenes[activeSceneIndex].StartingPositionX)
			tux.positionY = float64(scenes[activeSceneIndex].StartingPositionY)
			return
		}
	}
}
