package game

import (
	"game_of_life/pkg/board"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	g := New(10, 10)
	assert.Equal(t, g.Board.Ticks, 0)
	assert.Equal(t, 10, len(g.Board.Rows))
	memBoard := g.Board.Rows
	g.Tick()
	assert.NotEqual(t, memBoard, g.Board.Rows)
	assert.Equal(t, 1, g.Board.Ticks)
}

func TestPixDraw(t *testing.T) {
	r := [][]bool{
		{true, true, false},
		{false, false, true},
	}
	b := board.Board{Rows: r}
	g := GameWrapper{Board: b}
	inPix := make([]byte, 3*2*4)
	g.Draw(inPix)
	pix := make([]byte, 3*2*4)
	pix[0] = 0xff
	pix[1] = 0xff
	pix[2] = 0xff
	pix[3] = 0xff
	pix[4] = 0xff
	pix[5] = 0xff
	pix[6] = 0xff
	pix[7] = 0xff
	pix[8] = 0
	pix[9] = 0
	pix[10] = 0
	pix[11] = 0
	pix[12] = 0
	pix[13] = 0
	pix[14] = 0
	pix[15] = 0
	pix[16] = 0
	pix[17] = 0
	pix[18] = 0
	pix[19] = 0
	pix[20] = 0xff
	pix[21] = 0xff
	pix[22] = 0xff
	pix[23] = 0xff
	assert.Equal(t, pix, inPix)
}
