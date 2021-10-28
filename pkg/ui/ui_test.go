package ui

import (
	"game_of_life/pkg/game"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	g := game.GameWrapper{}
	u := New(&g)
	u.Update()
	assert.Equal(t, g.Board.Ticks, 1)
}

func TestLayout(t *testing.T) {
	g := game.GameWrapper{}
	u := New(&g)
	h, w := u.Layout(1, 2)
	assert.Equal(t, g.Height, h)
	assert.Equal(t, g.Width, w)
}

func TestDraw(t *testing.T) {
	g := game.GameWrapper{Height: 200, Width: 200}
	u := New(&g)
	s := ebiten.NewImage(g.Width, g.Height)
	u.Draw(s)
	assert.NotZero(t, len(g.Pixels))
}
