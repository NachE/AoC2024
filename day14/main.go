// Copyright J.A. Nache. MIT license
package main

import (
	"bufio"
	"fmt"
	"os"
)

var WIDTH int = 101
var HEIGHT int = 103

// always positive
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

func cpos(pp, vv [2]int) [2]int {
	npy := pos(pp[1], 100, vv[1], HEIGHT)
	npx := pos(pp[0], 100, vv[0], WIDTH)
	return [2]int{npx, npy}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var px, py, vx, vy, q1, q2, q3, q4 int
	// load entire file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Sscanf(string(scanner.Bytes()), "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		r := cpos([2]int{px, py}, [2]int{vx, vy})
		if r[0] == WIDTH/2 || r[1] == HEIGHT/2 {
			continue
		}

		switch {
		case r[0] > WIDTH/2 && r[1] < HEIGHT/2:
			q1++
		case r[0] < WIDTH/2 && r[1] < HEIGHT/2:
			q2++
		case r[0] < WIDTH/2 && r[1] > HEIGHT/2:
			q3++
		case r[0] > WIDTH/2 && r[1] > HEIGHT/2:
			q4++
		}
	}
	fmt.Println("res1", q1*q2*q3*q4)
}
