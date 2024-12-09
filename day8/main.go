// Copyright J.A. Nache. MIT license
package main

import (
	"bufio"
	"fmt"
	"os"
)

type board struct {
	grid       [][]byte
	rows       int
	cols       int
	dirs       map[[2]int]bool
	antinodes  map[[2]int]bool
	extranodes map[[2]int]bool
}

// get char at pos [2]int{row, col}
func (b *board) gc(pos [2]int) byte {
	return b.grid[pos[0]][pos[1]]
}

// set antinode, ignores out of greed
func (b *board) setAntinode(pos [2]int) {
	if isoutofgrid(b, pos) {
		return
	}
	b.antinodes[pos] = true
}

// set extra nodes, ignores out of greed
func (b *board) setExtraAntinodes(startPos [2]int, dir [2]int, step int) {
	b.extranodes[startPos] = true
	irow := startPos[0]
	icol := startPos[1]
	for {
		irow = irow + (dir[0] * step)
		icol = icol + (dir[1] * step)
		newantinode := [2]int{irow, icol}
		if isoutofgrid(b, newantinode) {
			break
		}
		b.extranodes[newantinode] = true
	}
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
			// because it is the same move with
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
// store antinodes cords in bb.antinodes
func search(bb *board, irow, icol int, dir [2]int) {
	step := 0
	f := bb.gc([2]int{irow, icol})
	for {
		step++
		nr := irow + (dir[0] * step)
		nc := icol + (dir[1] * step)
		if isoutofgrid(bb, [2]int{nr, nc}) {
			// out of grid
			return
		}
		if bb.gc([2]int{nr, nc}) == f {
			// tip for optimization: store antena pair in a map
			// to prevent re-check one pair twice
			// pair found
			// now should be an antinodeA at distance 'step' in direction 'dir'
			// and antinodeB starting from irow, icol in 'dir' rotated 180

			// look for antinode A for part 1
			antinodeA := [2]int{nr + (dir[0] * step), nc + (dir[1] * step)}
			bb.setAntinode(antinodeA)

			// calc antinodes for part 2 starting at 'nr', 'nc' in direction 'dir'
			bb.setExtraAntinodes([2]int{nr, nc}, dir, step)

			// rotate 'dir' 180 for part 2
			rdir := [2]int{-dir[0], -dir[1]}

			// lok for antinode B for part 1
			antinodeB := [2]int{irow + (rdir[0] * step), icol + (rdir[1] * step)}
			bb.setAntinode(antinodeB)

			// calc andinodes for part 2 starting at 'irow', 'icol' in direction 'rdir'
			bb.setExtraAntinodes([2]int{irow, icol}, rdir, step)
			// the pair antena has been found, break loop
			break
		}
	}
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
	bb.extranodes = make(map[[2]int]bool)

	for irow := 0; irow < bb.rows; irow++ {
		// iter board cols of current row
		for icol := 0; icol < bb.cols; icol++ {
			if bb.grid[irow][icol] != '.' && bb.grid[irow][icol] != '*' {
				// antena
				for dir := range bb.dirs {
					search(bb, irow, icol, dir)
				}
			}
		}
	}
	fmt.Println("res1: ", len(bb.antinodes))
	fmt.Println("res2: ", len(bb.extranodes))
}
