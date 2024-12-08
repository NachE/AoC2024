// Copyright J.A. Nache. MIT license
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var DUP = [2]int{-1, 0}
var DRIGHT = [2]int{0, 1}
var DDOWN = [2]int{1, 0}
var DLEFT = [2]int{0, -1}

type board struct {
	grid    [][]byte
	rows    int
	cols    int
	guard   [2]int
	dir     [2]int
	visited map[string]bool
}

func anotateCurrentGuardState(bb *board) bool {
	k := fmt.Sprintf("%v-%v(%v/%v)", bb.guard[0], bb.guard[1], bb.dir[0], bb.dir[1])
	if _, exists := bb.visited[k]; exists {
		return true
	}
	bb.visited[k] = true
	return false
}

func patrol(bb *board) (int, bool) { // bool: true -> loop found
	res := 1
	for {
		// already visited position & direction, loop lookup for part2
		k := fmt.Sprintf("%v-%v(%v/%v)", bb.guard[0], bb.guard[1], bb.dir[0], bb.dir[1])
		if _, exists := bb.visited[k]; exists {
			return res, true // loop found, return
		}
		bb.visited[k] = true

		// step pos
		nr := bb.guard[0] + bb.dir[0] // next row
		nc := bb.guard[1] + bb.dir[1] // next col

		if (nr < 0 || nr >= bb.rows) || (nc < 0 || nc >= bb.cols) {
			// out of grid
			break
		}

		// lookup for obstacle, rotate if found
		if bb.grid[nr][nc] == '#' {
			// rotate
			bb.dir = [2]int{bb.dir[1], -bb.dir[0]} // [-1, 0] (DUP) -> [0, 1] (DRIGHT)
			continue
		}

		// no obstacle found, allowed walk, seek
		bb.guard[0] = nr
		bb.guard[1] = nc

		// part1, store pos if new
		k = fmt.Sprintf("%v-%v", nr, nc)
		if _, exists := bb.visited[k]; !exists {
			bb.visited[k] = true
			res++
		}
	}
	return res, false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bb := &board{visited: make(map[string]bool)}

	scanner := bufio.NewScanner(file)
	bb.rows = -1
	gf := false // guard found
	for scanner.Scan() {
		bb.rows++
		lin := append([]byte{}, scanner.Bytes()...)
		if !gf {
			if r := bytes.Index(lin, []byte("^")); r > -1 {
				gf = true
				bb.guard = [2]int{bb.rows, r}
				bb.dir = DUP
			} else if r := bytes.Index(lin, []byte(">")); r > -1 {
				gf = true
				bb.guard = [2]int{bb.rows, r}
				bb.dir = DRIGHT
			} else if r := bytes.Index(lin, []byte("v")); r > -1 {
				gf = true
				bb.guard = [2]int{bb.rows, r}
				bb.dir = DDOWN
			} else if r := bytes.Index(lin, []byte("<")); r > -1 {
				gf = true
				bb.guard = [2]int{bb.rows, r}
				bb.dir = DLEFT
			}
		}
		bb.grid = append(bb.grid, lin)
	}
	bb.rows++
	bb.cols = len(bb.grid[0])
	bb.visited[string(bb.guard[0])+string(bb.guard[1])] = true
	startGuard := [2]int{bb.guard[0], bb.guard[1]}
	startDir := [2]int{bb.dir[0], bb.dir[1]}

	// patrol for part1
	res1, _ := patrol(bb)

	// test obstacles in every pos
	res2 := 0
	for irow := 0; irow < bb.rows; irow++ {
		// iter board cols of current row
		for icol := 0; icol < bb.cols; icol++ {
			// skip non-free position
			if bb.grid[irow][icol] == '#' {
				continue
			}
			// skip guard position
			if irow == startGuard[0] && icol == startGuard[1] {
				continue
			}
			// reset objects
			bb.guard = [2]int{startGuard[0], startGuard[1]}
			bb.dir = [2]int{startDir[0], startDir[1]}
			bb.visited = make(map[string]bool)
			bb.grid[irow][icol] = '#' // put tmp obstacle
			// patrol for loop
			if _, loop := patrol(bb); loop {
				// loop found
				res2++
			}
			bb.grid[irow][icol] = '.' // rm tmp obstacle
		}
	}

	fmt.Println("res1:", res1)
	fmt.Println("res2:", res2)
}
