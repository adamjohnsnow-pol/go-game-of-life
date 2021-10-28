package ui

import (
	"game_of_life/pkg/game"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type View struct {
	GameOfLife *game.GameWrapper
	Width      int
	Height     int
}

type Game interface {
	Draw(screen *ebiten.Image)
	Layout(int, int) (int, int)
	Update() error
}

func New(g *game.GameWrapper) Game {
	return &View{
		GameOfLife: g,
		Width:      g.Width,
		Height:     g.Height,
	}
}

func (v *View) Draw(screen *ebiten.Image) {
	if v.GameOfLife.Pixels == nil {
		v.GameOfLife.Pixels = make([]byte, v.Width*v.Height*4)
	}
	v.GameOfLife.Draw(v.GameOfLife.Pixels)
	screen.ReplacePixels(v.GameOfLife.Pixels)
}

func (v *View) Layout(int, int) (int, int) {
	return v.Width, v.Height
}

func (v *View) Update() error {
	v.GameOfLife.Tick()
	time.Sleep(time.Second / 26)
	return nil
}
