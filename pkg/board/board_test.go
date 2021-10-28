package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	b, err := New(5, 5)
	assert.NoError(t, err)
	assert.Equal(t, 0, b.Ticks)
	assert.Len(t, b.Rows, 5)
}

func TestRandom(t *testing.T) {
	row := randomRow(5, true)
	assert.Len(t, row, 5)
	row2 := randomRow(5, true)
	assert.Equal(t, row, row2)
	row2 = randomRow(5, false)
	assert.NotEqual(t, row, row2)
}

func TestIncrementTick(t *testing.T) {
	r := [][]bool{
		{true, false, true, true},
		{true, true, true, true},
		{false, false, true, true},
		{false, false, true, true},
	}
	b := Board{Rows: r}

	b.IncrementTick()
	assert.Equal(
		t,
		[][]bool{
			{true, false, false, true},
			{true, false, false, false},
			{false, false, false, false},
			{false, false, true, true},
		},
		b.Rows,
	)
	assert.Equal(t, b.Ticks, 1)
	b.IncrementTick()
	assert.Equal(
		t,
		[][]bool{
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
		},
		b.Rows,
	)
	assert.Equal(t, b.Ticks, 2)

	r = [][]bool{
		{true, false, true, true, false, false, false},
		{true, true, true, true, false, false, false},
		{false, false, true, true, false, false, false},
		{false, false, true, true, false, false, false},
	}
	b = Board{Rows: r}

	b.IncrementTick()
	assert.Equal(t, []bool{true, false, false, true, false, false, false}, b.Rows[0])
	assert.Equal(t, []bool{true, false, false, false, true, false, false}, b.Rows[1])
	assert.Equal(t, []bool{false, false, false, false, true, false, false}, b.Rows[2])
	assert.Equal(t, []bool{false, false, true, true, false, false, false}, b.Rows[3])

	b.IncrementTick()
	assert.Equal(t, []bool{false, false, false, false, false, false, false}, b.Rows[0])
	assert.Equal(t, []bool{false, false, false, true, true, false, false}, b.Rows[1])
	assert.Equal(t, []bool{false, false, false, false, true, false, false}, b.Rows[2])

}

func TestEvaluateNeighbours(t *testing.T) {
	cell := [3][3]bool{
		{true, false, true},
		{true, false, true},
		{true, false, true},
	}
	r := evaluateNeighbours(cell)
	assert.False(t, r)

	cell = [3][3]bool{
		{true, false, true},
		{false, false, false},
		{false, false, true},
	}
	r = evaluateNeighbours(cell)
	assert.True(t, r)

	cell = [3][3]bool{
		{true, false, false},
		{false, false, false},
		{false, true, true},
	}
	r = evaluateNeighbours(cell)
	assert.True(t, r)

	cell = [3][3]bool{
		{true, false, false},
		{false, true, false},
		{false, false, true},
	}
	r = evaluateNeighbours(cell)
	assert.True(t, r)

	cell = [3][3]bool{
		{false, false, false},
		{false, true, false},
		{false, false, false},
	}
	r = evaluateNeighbours(cell)
	assert.False(t, r)

}

func TestCountTrue(t *testing.T) {
	i := countTrue([3]bool{true, false, true})
	assert.Equal(t, 2, i)
	i = countTrue([3]bool{true, true, true})
	assert.Equal(t, 3, i)
}

func TestGetCell(t *testing.T) {
	r := [][]bool{
		{true, false, true, true},
		{true, true, true, true},
		{false, false, true, true},
		{false, false, true, true},
	}
	b := Board{Rows: r}
	cell := b.getCell(1, 1)
	assert.Equal(t, [3]bool{true, false, true}, cell[0])
	assert.Equal(t, [3]bool{true, true, true}, cell[1])
	assert.Equal(t, [3]bool{false, false, true}, cell[2])

	cell = b.getCell(0, 0)
	assert.Equal(t, [3]bool{false, false, false}, cell[0])
	assert.Equal(t, [3]bool{false, true, false}, cell[1])
	assert.Equal(t, [3]bool{false, true, true}, cell[2])

	cell = b.getCell(3, 3)
	assert.Equal(t, [3]bool{true, true, false}, cell[0])
	assert.Equal(t, [3]bool{true, true, false}, cell[1])
	assert.Equal(t, [3]bool{false, false, false}, cell[2])

}

func TestStablePattern(t *testing.T) {
	r := [][]bool{
		{false, false, false, false},
		{false, true, true, false},
		{false, true, true, false},
		{false, false, false, false},
	}
	b := Board{Rows: r}
	b.IncrementTick()
	b.IncrementTick()
	b.IncrementTick()
	b.IncrementTick()
	b.IncrementTick()
	b.IncrementTick()
	assert.Equal(t, b.Rows, r)

	r = [][]bool{
		{false, true, true, false},
		{true, false, false, true},
		{false, true, true, false},
	}
	b = Board{Rows: r}
	b.IncrementTick()
	b.IncrementTick()
	b.IncrementTick()
	b.IncrementTick()
	b.IncrementTick()
	b.IncrementTick()
	assert.Equal(t, b.Rows, r)

	//glider - after 4 ticks should transpose same shape 1 up, 1 left
	r = [][]bool{
		{false, false, false, false, false},
		{false, false, true, true, true},
		{false, false, true, false, false},
		{false, false, false, true, false},
	}
	b = Board{Rows: r}
	b.IncrementTick()
	b.IncrementTick()
	b.IncrementTick()
	b.IncrementTick()
	ex := [][]bool{
		{false, true, true, true, false},
		{false, true, false, false, false},
		{false, false, true, false, false},
		{false, false, false, false, false},
	}
	assert.Equal(t, ex[0], b.Rows[0])
	assert.Equal(t, ex[1], b.Rows[1])
	assert.Equal(t, ex[2], b.Rows[2])
	assert.Equal(t, ex[3], b.Rows[3])
}
