package board

import (
	"math/rand"
	"time"
)

type Board struct {
	Ticks       int
	Rows        [][]bool
	ActiveCount int
}

func New(y, x int) (Board, error) {
	b := Board{
		Ticks: 0,
	}
	for i := 0; i < y; i++ {
		b.Rows = append(b.Rows, randomRow(x, false))
	}

	return b, nil
}

func randomRow(x int, fix bool) []bool {
	row := make([]bool, x)
	for i := range row {
		row[i] = randBool(fix)
	}
	return row
}

func randBool(fix bool) bool {
	if fix {
		rand.Seed(123)
	} else {
		rand.Seed(time.Now().UnixNano())
	}
	return rand.Intn(2) == 1
}

func (b *Board) IncrementTick() {
	b.Ticks = b.Ticks + 1
	b.ActiveCount = 0
	var cell [3][3]bool
	var newBoard [][]bool
	for i, row := range b.Rows {
		var newRow []bool
		for c := range row {
			cell = b.getCell(i, c)
			newState := evaluateNeighbours(cell)
			if newState {
				b.ActiveCount++
			}
			newRow = append(newRow, newState)
		}
		newBoard = append(newBoard, newRow)
	}
	b.Rows = newBoard
}

func (b *Board) getCell(y, x int) [3][3]bool {
	var cell [3][3]bool
	h := len(b.Rows) - 1
	w := len(b.Rows[0]) - 1
	cell[1][1] = b.Rows[y][x]

	if x > 0 && y > 0 {
		cell[0][0] = b.Rows[y-1][x-1]
	}
	if y > 0 {
		cell[0][1] = b.Rows[y-1][x]
	}
	if y > 0 && x < w {
		cell[0][2] = b.Rows[y-1][x+1]
	}
	if x > 0 {
		cell[1][0] = b.Rows[y][x-1]
	}
	if x < w {
		cell[1][2] = b.Rows[y][x+1]
	}
	if x > 0 && y < h {
		cell[2][0] = b.Rows[y+1][x-1]
	}
	if y < h {
		cell[2][1] = b.Rows[y+1][x]
	}
	if x < w && y < h {
		cell[2][2] = b.Rows[y+1][x+1]
	}
	return cell
}

func evaluateNeighbours(cell [3][3]bool) bool {
	topRow := countTrue(cell[0])
	midRow := countTrue([3]bool{cell[1][0], false, cell[1][2]})
	bottomRow := countTrue(cell[2])
	sum := topRow + midRow + bottomRow
	if sum == 3 {
		return true
	}
	if sum == 2 && cell[1][1] {
		return true
	}
	return false
}

func countTrue(b [3]bool) int {
	n := 0
	for _, v := range b {
		if v {
			n++
		}
	}
	return n
}
