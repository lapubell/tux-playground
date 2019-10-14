package main

import (
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lapubell/tux-playground/sprite"
	"golang.org/x/image/colornames"
)

func app() {
	cfg := pixelgl.WindowConfig{
		Title:  "Tux's Playground",
		Bounds: pixel.R(0, 0, 800, 600),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	t := sprite.NewTux()

	// last := time.Now()

	for !win.Closed() {
		win.Clear(colornames.Grey)

		// dt := time.Since(last).Seconds()
		// last = time.Now()
		mat := pixel.IM.Moved(pixel.V(10, 10))
		t.Draw(win, mat)

		if win.Pressed(pixelgl.KeyLeft) {
			t.UpdateXAcceleration(-0.1)
		} else {
			t.Decelerate()
		}

		if win.Pressed(pixelgl.KeyRight) {
			t.UpdateXAcceleration(0.1)
		} else {
			t.Decelerate()
		}

		if win.Pressed(pixelgl.KeyUp) {
			t.UpdateYAcceleration(0.1)
		} else {
			t.Decelerate()
		}

		if win.Pressed(pixelgl.KeyDown) {
			t.UpdateYAcceleration(-0.1)
		} else {
			t.Decelerate()
		}

		win.Update()
	}
}
