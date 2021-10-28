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

func New(y, x int, random bool) (Board, error) {
	b := Board{
		Ticks: 0,
	}
	if random {
		for i := 0; i < y; i++ {
			b.Rows = append(b.Rows, randomRow(x, false))
		}
	} else {
		b.Rows = getGliderGun(x, y)
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

func getGliderGun(x, y int) [][]bool {
	var rows [][]bool
	for i := 0; i < y; i++ {
		rows = append(rows, make([]bool, x))
	}
	coords := [][]int{
		{1, 25},
		{2, 23}, {2, 25},
		{3, 13}, {3, 14}, {3, 21}, {3, 22}, {3, 35}, {3, 36},
		{4, 12}, {4, 16}, {4, 21}, {4, 22}, {4, 35}, {4, 36},
		{5, 1}, {5, 2}, {5, 11}, {5, 17}, {5, 21}, {5, 22},
		{6, 1}, {6, 2}, {6, 11}, {6, 15}, {6, 17}, {6, 18}, {6, 23}, {6, 25},
		{7, 11}, {7, 17}, {7, 25},
		{8, 12}, {8, 16},
		{9, 13}, {9, 14},
	}
	for _, c := range coords {
		rows[c[0]][c[1]] = true
	}
	return rows
}
