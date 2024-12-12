// Copyright J.A. Nache. MIT license
package main

import (
	"bufio"
	"fmt"
	"os"
)

var dirs [4][2]int = [4][2]int{
	{0, 1},  // right
	{-1, 0}, // up
	{0, -1}, // left
	{1, 0},  // down
}

type state struct {
	plant byte
	pos   [2]int
}

type board struct {
	grid    [][]byte
	visited map[[2]int]bool
	rows    int
	cols    int
}

func rotate__90(dir [2]int) [2]int {
	return [2]int{-dir[1], dir[0]}
}

func rotate90(dir [2]int) [2]int {
	return [2]int{dir[1], -dir[0]}
}

// Return position after rotating -90 and moving 1 step
func rotate__90step(pos [2]int, dir [2]int) [2]int {
	dir = rotate__90(dir) // rotate -90
	return [2]int{pos[0] + dir[0], pos[1] + dir[1]}
}

// Get the character at position [2]int{row, col}
// Return true if it is out of bounds
func (b *board) get(pos [2]int) (bool, byte) {
	if (pos[0] < 0 || pos[0] >= b.rows) || (pos[1] < 0 || pos[1] >= b.cols) {
		return true, '*'
	}
	return false, b.grid[pos[0]][pos[1]]
}

func (b *board) isvisited(pos [2]int) bool {
	_, exists := b.visited[pos]
	return exists
}

func (b *board) visite(pos [2]int) {
	b.visited[pos] = true
}

func (b *board) isoutside(pos [2]int, plant byte) bool {
	out, _p := b.get(pos)
	if out || _p != plant {
		return true // wall found
	}
	return false
}

// A wall-following algoritm
func (b *board) maze(startpos [2]int, outsides map[[2]int]bool, startdir [2]int) int {
	dir := startdir
	plant := b.grid[startpos[0]][startpos[1]]
	pos := startpos
	sides := 0

	// Test a single field trap
	_i := 0
	for _, _d := range dirs {
		rpos := rotate__90step(pos, _d)
		if b.isoutside(rpos, plant) {
			delete(outsides, rpos)
			_i++
		}
	}
	if _i == 4 {
		return 4
	}

	for {
		// Warning: Assumes the starting position is always on a wall.
		// No initial wall lookup is needed.
		// This strategic position should be pre-calculated
		// before calling maze().
		rpos := rotate__90step(pos, dir)
		if !b.isoutside(rpos, plant) {
			// No wall found, continue rotating and walking
			dir = rotate__90(dir)
			pos = rpos
			sides++
		} else {
			// Wall found
			delete(outsides, rpos)
			nextmove := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
			// Walk 1 pos
			out, _p := b.get(nextmove)
			if out || _p != plant {
				dir = rotate90(dir)
				sides++
			} else {
				pos = nextmove
			}
		}

		// Loop finished, stop walking
		if startpos[0] == pos[0] && startpos[1] == pos[1] &&
			dir[0] == startdir[0] && dir[1] == startdir[1] {
			break
		}
	}
	return sides
}

// DFS + wall-following
func (b *board) calcregion(startpos [2]int) (int, int) {
	area := 1 // 1 -> self
	perim := 0

	plant := b.grid[startpos[0]][startpos[1]]
	stck := []state{{
		plant: plant,
		pos:   startpos,
	}}
	b.visite(startpos)
	outsides := make(map[[2]int]bool)

	for len(stck) > 0 {
		stte := stck[len(stck)-1]
		stck = stck[:len(stck)-1]

		for _, dir := range dirs {
			pos := [2]int{stte.pos[0] + dir[0], stte.pos[1] + dir[1]}
			out, plant := b.get(pos)
			if out {
				outsides[pos] = true
				perim++
				continue
			}
			if plant != stte.plant {
				outsides[pos] = true
				perim++

			} else {
				if b.isvisited(pos) {
					continue
				}
				area++
				b.visite(pos)
				stck = append(stck, state{
					plant: plant,
					pos:   pos,
				})
			}
		}
	}

	// Calculate external sides: the starting position is at
	// an external wall
	sides := b.maze(startpos, outsides, dirs[0])
	// Calculate internal sides while there are outsides
	for len(outsides) > 0 {
		for op := range outsides {
			// Look for an outside edge that will keep the
			// starting position at the internal wall
			rr := [2]int{op[0], op[1] + 1}
			out, _p := b.get(rr)
			if !out && _p == plant {
				sides += b.maze(rr, outsides, dirs[1])
				break
			}
		}
	}

	return area * perim, area * sides
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bb := &board{}

	// Load entire file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lin := append([]byte{}, scanner.Bytes()...)
		bb.grid = append(bb.grid, lin)
	}
	bb.rows = len(bb.grid)
	bb.cols = len(bb.grid[0])
	bb.visited = make(map[[2]int]bool)

	res1 := 0
	res2 := 0
	for irow := 0; irow < bb.rows; irow++ {
		for icol := 0; icol < bb.cols; icol++ {
			if !bb.isvisited([2]int{irow, icol}) {
				perimc, sidesc := bb.calcregion([2]int{irow, icol})
				res1 += perimc
				res2 += sidesc
			}
		}
	}

	fmt.Println("res1:", res1)
	fmt.Println("res2:", res2)
}
