package game

import (
	"game_of_life/pkg/board"
)

type GameWrapper struct {
	Height int
	Width  int
	Board  board.Board
	Pixels []byte
}

func New(w, h int) GameWrapper {
	b, err := board.New(h, w)
	if err != nil {
		panic(err)
	}
	return GameWrapper{
		Height: h,
		Width:  w,
		Board:  b,
		Pixels: nil}
}

func (g *GameWrapper) Tick() {
	g.Board.IncrementTick()
}

func (g *GameWrapper) Draw(pix []byte) {
	p := 0
	for _, r := range g.Board.Rows {
		for _, v := range r {
			if v {
				pix[p] = 0xff
				pix[p+1] = 0xff
				pix[p+2] = 0xff
				pix[p+3] = 0xff
			} else {
				pix[p] = 0
				pix[p+1] = 0
				pix[p+2] = 0
				pix[p+3] = 0
			}
			p = p + 4
		}
	}
}
