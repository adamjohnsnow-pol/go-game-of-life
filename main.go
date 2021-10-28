package main

import (
	"game_of_life/pkg/game"
	"game_of_life/pkg/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 300
	screenHeight = 250
)

func main() {
	g := game.New(screenWidth, screenHeight)
	ui := ui.New(g)
	makeGame(ui)
}

func makeGame(ui ui.Game) {
	ebiten.SetWindowSize(screenWidth*4, screenHeight*4)
	ebiten.SetWindowTitle("Go Game of Life")
	ebiten.RunGame(ui)
}
