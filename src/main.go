package main

import (
	"log"

	"github.com/lapubell/tux-playground/game"
)

func main() {
	g := game.Game{}
	err := g.Run()
	if err != nil {
		log.Fatal(err)
	}
}
