package main

import (
	"game_of_life/pkg/game"
	"game_of_life/pkg/ui"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 120
	screenHeight = 100
)

func main() {
	random := len(os.Args) == 1
	g := game.New(screenWidth, screenHeight, random)
	ui := ui.New(g)
	makeGame(ui)
}

func makeGame(ui ui.Game) {
	ebiten.SetWindowSize(screenWidth*6, screenHeight*6)
	ebiten.SetWindowTitle("Go Game of Life")
	ebiten.RunGame(ui)
}
