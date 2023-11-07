package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 640
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Generic Shooting Game")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
