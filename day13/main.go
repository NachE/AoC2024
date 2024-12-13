// Copyright J.A. Nache. MIT license
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var Xaxis = 0
var Yaxis = 1

type machineBehavior struct {
	buttonA [2]int
	buttonB [2]int
	prize   [2]int
}

func (m *machineBehavior) parseLine(bline []byte, split byte) (int, int) {
	btn := bytes.Split(bline, []byte{split})
	x, err := strconv.Atoi(string(bytes.TrimRight(btn[1], ", Y")))
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(string(btn[2]))
	if err != nil {
		panic(err)
	}
	return x, y
}

func (m *machineBehavior) parseButtonA(bline []byte) {
	x, y := m.parseLine(bline, '+')
	m.buttonA = [2]int{x, y}
}

func (m *machineBehavior) parseButtonB(bline []byte) {
	x, y := m.parseLine(bline, '+')
	m.buttonB = [2]int{x, y}
}

func (m *machineBehavior) parsePrice(bline []byte) {
	x, y := m.parseLine(bline, '=')
	m.prize = [2]int{x, y}
}

func gcf(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Find the best solution for the machine using its behavior (mb)
func play(mb *machineBehavior) int {
	// X axis
	gcfX := gcf(mb.buttonA[Xaxis], mb.buttonB[Xaxis])
	if mb.prize[Xaxis]%gcfX != 0 {
		// No solution in X axis
		return 0
	}

	// Y axis
	gcfY := gcf(mb.buttonA[Yaxis], mb.buttonB[Yaxis])
	if mb.prize[Yaxis]%gcfY != 0 {
		// No solution in Y axis
		return 0
	}

	cost := 0
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			// Test prize values from 0 to 100
			// (max 100: "You estimate that each button would need
			// to be pressed no more than 100 times")
			if a*mb.buttonA[Xaxis]+b*mb.buttonB[Xaxis] == mb.prize[Xaxis] &&
				a*mb.buttonA[Yaxis]+b*mb.buttonB[Yaxis] == mb.prize[Yaxis] {
				cost = 3*a + b
			}
		}
	}

	return cost
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res1 := 0
	// load entire file
	scanner := bufio.NewScanner(file)
	for {
		mb := &machineBehavior{}
		scanner.Scan()
		mb.parseButtonA(append([]byte{}, scanner.Bytes()...))

		scanner.Scan()
		mb.parseButtonB(append([]byte{}, scanner.Bytes()...))

		scanner.Scan()
		mb.parsePrice(append([]byte{}, scanner.Bytes()...))

		el := scanner.Scan() // empty line

		res1 += play(mb)

		if !el {
			break
		}
	}

	fmt.Println("res1:", res1)
}
