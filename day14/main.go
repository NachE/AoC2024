// Copyright J.A. Nache. MIT license
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var WIDTH int = 101
var HEIGHT int = 103
var X = 0
var Y = 1

type robot struct {
	X int
	Y int
}

func (r *robot) int2pp() [2]int {
	return [2]int{r.X, r.Y}
}

// Always positive
// useful when working with circular
// spaces (such as list indices or
// wrapped coordinates).
func mod(a, b int) int {
	return (a%b + b) % b
}

// xy := | (p + t * v) % l |
// p -> initial position
// t -> seconds
// v -> speed (in 1 second)
// l -> width or height}
func pos(p, s, v, l int) int {
	return mod(p+s*v, l)
}

// For part 1, calculate the final position from
// the initial pp position and vv velocity.
func cpos(pp, vv [2]int, tt int) [2]int {
	npy := pos(pp[Y], tt, vv[Y], HEIGHT)
	npx := pos(pp[X], tt, vv[X], WIDTH)
	return [2]int{npx, npy}
}

// Return true if there is any robot at position rr at time tt.
func forwardrobots(robots map[*robot][2]int, rr [2]int, tt int) bool {
	for robot, vv := range robots {
		r := cpos(robot.int2pp(), vv, tt)
		if r[X] == rr[X] && r[Y] == rr[Y] {
			return true
		}
	}
	return false
}

// Return true if a robot starting at position pp with velocity vv
// can reach position rr. Also returns the time it takes.
func rcpos(pp, rr, vv [2]int, start, limit int) (bool, int) {
	for t := start; t < limit; t++ {
		x := mod(pp[X]+t*vv[X], WIDTH)
		y := mod(pp[X]+t*vv[Y], HEIGHT)

		if x == rr[0] && y == rr[1] {
			return true, t
		}
	}
	return false, -1
}

// Calculate for Part 1
func quarts(robots map[*robot][2]int, tt int) (int, int, int, int) {
	var q1, q2, q3, q4 int
	for robot, vv := range robots {
		r := cpos(robot.int2pp(), vv, tt)
		if r[X] == WIDTH/2 || r[Y] == HEIGHT/2 {
			continue
		}

		switch {
		case r[X] > WIDTH/2 && r[Y] < HEIGHT/2:
			q1++
		case r[X] < WIDTH/2 && r[Y] < HEIGHT/2:
			q2++
		case r[X] < WIDTH/2 && r[Y] > HEIGHT/2:
			q3++
		case r[X] > WIDTH/2 && r[Y] > HEIGHT/2:
			q4++
		}
	}
	return q1, q2, q3, q4
}

// Prints the board status at time tt
func printboard(robots map[*robot][2]int, tt int) {
	var board [][]byte = make([][]byte, 0)
	for i := 0; i < HEIGHT; i++ {
		board = append(board, bytes.Repeat([]byte("."), WIDTH))
	}
	for robot, vv := range robots {
		r := cpos(robot.int2pp(), vv, tt)
		board[r[Y]][r[X]] = '#'
	}
	for i := 0; i < len(board); i++ {
		fmt.Println(string(board[i]))
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	robots := make(map[*robot][2]int)

	var px, py, vx, vy int
	// load entire file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Sscanf(string(scanner.Bytes()), "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		rob := &robot{X: px, Y: py}
		robots[rob] = [2]int{vx, vy}
	}

	q1, q2, q3, q4 := quarts(robots, 100)
	fmt.Println("res1", q1*q2*q3*q4)

	res2 := -1
	rr := [2]int{(WIDTH / 2), (HEIGHT / 2)}
	l := WIDTH * HEIGHT
	s := 0
L:
	for {
		for robot, vv := range robots {
			f, t := rcpos(robot.int2pp(), rr, vv, s, l)
			if !f {
				continue
			}
			// A small hack here: look for a cross in the center ;)
			right := [2]int{rr[X] + 1, rr[Y]}
			if !forwardrobots(robots, right, t) {
				continue
			}
			left := [2]int{rr[X] - 1, rr[Y]}
			if !forwardrobots(robots, left, t) {
				continue
			}
			up := [2]int{rr[X], rr[Y] - 1}
			if !forwardrobots(robots, up, t) {
				continue
			}
			down := [2]int{rr[X], rr[Y] + 1}
			if !forwardrobots(robots, down, t) {
				continue
			}
			printboard(robots, t)
			res2 = t
			break L
		}
		s = l - 1
		l = l * 4
	}

	fmt.Println("res2", res2)
}
