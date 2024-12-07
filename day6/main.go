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
		// .Bytes() returns a referente to internal buff
		// so we need to copy it in order to prevent
		// multi-referenced board wid dupe data
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
	bb.cols = len(bb.grid[0])
	bb.visited[string(bb.guard[0])+string(bb.guard[1])] = true

	res1 := 1
	for {
		nr := bb.guard[0] + bb.dir[0] // next row
		nc := bb.guard[1] + bb.dir[1] // next col

		if nr < 0 || nr >= bb.rows {
			// out of board
			break
		}
		if nc < 0 || nc >= bb.cols {
			// out of board
			break
		}
		if bb.grid[nr][nc] == '#' {
			// rotate
			bb.dir = [2]int{bb.dir[1], -bb.dir[0]} // [-1, 0] (DUP) -> [0, 1] (DRIGHT)
			continue
		}
		bb.guard[0] = nr
		bb.guard[1] = nc

		if _, exists := bb.visited[string(nr)+string(nc)]; !exists {
			bb.visited[string(nr)+string(nc)] = true
			res1++
			// dbg:
			// bb.grid[nr][nc] = '*'
		}
	}
	// dbg:
	// for _, row := range bb.grid {
	// 	fmt.Println(string(row))
	// }

	fmt.Println("res1:", res1)
}
