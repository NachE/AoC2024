// Copyright J.A. Nache. MIT license
package main

import (
	"bufio"
	"fmt"
	"os"
)

type board struct {
	grid      [][]byte
	rows      int
	cols      int
	dirs      map[[2]int]bool
	antinodes map[[2]int]bool
}

// get char at pos [2]int{row, col}
func (b *board) gc(pos [2]int) byte {
	return b.grid[pos[0]][pos[1]]
}

// set antinode, return false if already exists
// or if the pos is out of map and antinode
// cannot be setted
func (b *board) setAntinode(pos [2]int) bool {
	if isoutofgrid(b, pos) {
		return false
	}
	if _, exists := b.antinodes[pos]; exists {
		return false // already setted antinode
	}
	b.antinodes[pos] = true
	return true // new antinode
}

// absolute greatest common factor
func agcf(a, b int) int {
	if a < 0 {
		a = -a // rm sign
	}
	if b < 0 {
		b = -b // rm sign
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// build possible vectors of directions, radar like
// dx, dy == {-r, ..., r}
// dcol -> dx (distance in x (col))
// drow -> dy (distance in y (row))
func builddirs(r int) map[[2]int]bool {
	dirs := make(map[[2]int]bool)

	for drow := -r; drow <= r; drow++ {
		for dcol := -r; dcol <= r; dcol++ {
			if drow == 0 && dcol == 0 {
				continue // skip no move
			}

			// prevent -1, -1 and -2, -2 and -4, -4
			// because is the same move but with
			// different rad size
			agcf := agcf(drow, dcol)
			ndrow := drow / agcf
			ndcol := dcol / agcf

			// store in a map to prevent dupes
			dirs[[2]int{ndrow, ndcol}] = true
		}
	}

	return dirs
}

// cords [2]int{row, col}
func isoutofgrid(bb *board, cords [2]int) bool {
	if (cords[0] < 0 || cords[0] >= bb.rows) || (cords[1] < 0 || cords[1] >= bb.cols) {
		// out of grid
		return true
	}
	return false
}

// search for antinodes in bb starting at irow, icol in dir direction
// store antinodes cords in bb.antinodes for no dupes-track
// return num of antinodes found
func search(bb *board, irow, icol int, dir [2]int) int {
	step := 0
	f := bb.gc([2]int{irow, icol})
	r := 0
	for {
		step++
		nr := irow + (dir[0] * step)
		nc := icol + (dir[1] * step)
		if isoutofgrid(bb, [2]int{nr, nc}) {
			// out of grid
			return 0
		}
		if bb.gc([2]int{nr, nc}) == f {
			// tip for optimization: store antena pair in a map
			// to prevent re-check one pair twice
			// pair found
			// now should be a antinodeA at distance step with in dir
			// and antinodeB starting from irow, icol with dir rotated 180
			antinodeA := [2]int{nr + (dir[0] * step), nc + (dir[1] * step)}
			if bb.setAntinode(antinodeA) {
				r++
			}
			// rotate dir 180
			rdir := [2]int{-dir[0], -dir[1]}
			antinodeB := [2]int{irow + (rdir[0] * step), icol + (rdir[1] * step)}
			if bb.setAntinode(antinodeB) {
				r++
			}
			break // the pair antena has been found, break loop
		}
	}
	return r
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
	bb.dirs = builddirs(bb.rows * 2)
	bb.antinodes = make(map[[2]int]bool)

	res1 := 0
	for irow := 0; irow < bb.rows; irow++ {
		// iter board cols of current row
		for icol := 0; icol < bb.cols; icol++ {
			if bb.grid[irow][icol] != '.' && bb.grid[irow][icol] != '*' {
				// antena
				for dir := range bb.dirs {
					res1 += search(bb, irow, icol, dir)
				}
			}
		}
	}
	fmt.Println("res1: ", res1)
}
