// Copyright J.A. Nache. MIT license
package main

import (
	"bufio"
	"fmt"
	"os"
)

var dirs [4][2]int = [4][2]int{
	{0, 1},  // right
	{0, -1}, // left
	{-1, 0}, // up
	{1, 0},  // down
}

var steps map[byte]int = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

type state struct {
	pos  [2]int
	step int
}

type board struct {
	grid [][]byte
	rows int
	cols int
}

func (b *board) get(pos [2]int) int {
	if (pos[0] < 0 || pos[0] >= b.rows) || (pos[1] < 0 || pos[1] >= b.cols) {
		// out of board
		return -1
	}

	bn := b.grid[pos[0]][pos[1]]
	return steps[bn]
}

func (b *board) findpaths(startpos [2]int) int {
	score := 0

	stck := []state{{
		pos:  startpos,
		step: 0,
	}}

	found := make(map[[2]int]bool)

	for len(stck) > 0 {
		// pop one
		stte := stck[len(stck)-1]
		stck = stck[:len(stck)-1]

		for _, dir := range dirs {
			rowmove := stte.pos[0] + dir[0]
			colmove := stte.pos[1] + dir[1]

			if b.get([2]int{rowmove, colmove}) == stte.step+1 {
				if stte.step+1 == 9 {
					if _, exists := found[[2]int{rowmove, colmove}]; !exists {
						found[[2]int{rowmove, colmove}] = true
						score++
					}
					continue
				}
				// push next move
				stck = append(stck, state{
					pos:  [2]int{rowmove, colmove},
					step: stte.step + 1,
				})
			}
		}
	}

	return score
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bb := &board{}

	// load entire file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lin := append([]byte{}, scanner.Bytes()...)
		bb.grid = append(bb.grid, lin)
	}
	bb.rows = len(bb.grid)
	bb.cols = len(bb.grid[0])

	res1 := 0
	for irow := 0; irow < bb.rows; irow++ {
		for icol := 0; icol < bb.cols; icol++ {
			if bb.grid[irow][icol] == '0' {
				r := bb.findpaths([2]int{irow, icol})
				res1 += r
			}
		}
	}

	fmt.Println("res1:", res1)

}
