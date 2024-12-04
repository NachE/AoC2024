// Copyright J.A. Nache. MIT license
package main

import (
	"bufio"
	"fmt"
	"os"
)

// [2]int{row, col}

var RIGHT = [2]int{0, 1}
var LEFT = [2]int{0, -1}

var UP = [2]int{1, 0}
var DOWN = [2]int{-1, 0}

var D_DOWN_RIGHT = [2]int{1, 1}
var D_DOWN_LEFT = [2]int{1, -1}

var D_UP_LEFT = [2]int{-1, -1}
var D_UP_RIGHT = [2]int{-1, 1}

var moves [8][2]int = [8][2]int{
	RIGHT,
	LEFT,
	UP,
	DOWN,
	D_DOWN_RIGHT,
	D_UP_LEFT,
	D_DOWN_LEFT,
	D_UP_RIGHT,
}

type board struct {
	grid [][]byte
	rows int
	cols int
}

func search(b *board, irow, icol int, dir [2]int, word []byte) bool {
	for ichr := 0; ichr < len(word); ichr++ {
		rowmove := irow + (dir[0] * ichr)
		colmove := icol + (dir[1] * ichr)

		if rowmove < 0 || rowmove >= b.rows {
			// out of board
			return false
		}
		if colmove < 0 || colmove >= b.cols {
			// out of board
			return false
		}
		if b.grid[rowmove][colmove] != word[ichr] {
			// bad char
			return false
		}
	}
	return true
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
		// .Bytes() returns a referente to internal buff
		// so we need to copy it in order to prevent
		// multi-referenced board wid dupe data
		lin := append([]byte{}, scanner.Bytes()...)
		bb.grid = append(bb.grid, lin)
	}

	bb.rows = len(bb.grid)
	bb.cols = len(bb.grid[0])

	c := 0
	// part 1
	for irow := 0; irow < bb.rows; irow++ {
		// iter board cols of current row
		for icol := 0; icol < bb.cols; icol++ {
			// check char of current board cursor position
			iword := bb.grid[irow][icol]
			if iword != 'X' {
				continue // no X, skip
			}
			// here we found X of XMAS
			// search in all directions
			for _, dir := range moves {
				// search XMAS word
				if !search(bb, irow, icol, dir, []byte("XMAS")) {
					continue
				}
				c++
			}
		}
	}

	c2 := 0
	// part 2
	for irow := 0; irow < bb.rows; irow++ {
		// iter board cols of current row
		for icol := 0; icol < bb.cols; icol++ {
			// check char of current board cursor position
			iword := bb.grid[irow][icol]
			if iword != 'A' {
				continue // no A, skip
			}
			// here we found A of MAS
			// now look for MAS

			// for diagonal \ in both directions
			_a := search(bb, irow-1, icol-1, D_DOWN_RIGHT, []byte("MAS"))
			_b := search(bb, irow+1, icol+1, D_UP_LEFT, []byte("MAS"))
			if !_a && !_b {
				continue
			}

			// for diagonal / in both directions
			_a = search(bb, irow-1, icol+1, D_DOWN_LEFT, []byte("MAS"))
			_b = search(bb, irow+1, icol-1, D_UP_RIGHT, []byte("MAS"))
			if !_a && !_b {
				continue
			}

			c2++
		}
	}

	fmt.Println("res1:", c)
	fmt.Println("res2:", c2)
}
