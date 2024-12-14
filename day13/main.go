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
	buttonA   [2]int64
	buttonB   [2]int64
	prize     [2]int64
	prizebase [2]int64
}

func (m *machineBehavior) parseLine(bline []byte, split byte) (int64, int64) {
	btn := bytes.Split(bline, []byte{split})
	x, err := strconv.Atoi(string(bytes.TrimRight(btn[1], ", Y")))
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(string(btn[2]))
	if err != nil {
		panic(err)
	}
	return int64(x), int64(y)
}

func (m *machineBehavior) parseButtonA(bline []byte) {
	x, y := m.parseLine(bline, '+')
	m.buttonA = [2]int64{x, y}
}

func (m *machineBehavior) parseButtonB(bline []byte) {
	x, y := m.parseLine(bline, '+')
	m.buttonB = [2]int64{x, y}
}

func (m *machineBehavior) parsePrice(bline []byte) {
	x, y := m.parseLine(bline, '=')
	m.prize = [2]int64{x, y}
}

func play(mb *machineBehavior) int64 {
	przx := mb.prize[0] + mb.prizebase[0]
	przy := mb.prize[1] + mb.prizebase[1]

	// Calculate the determinant
	// https://en.wikipedia.org/wiki/Determinant
	dt := (mb.buttonA[0]*mb.buttonB[1] - mb.buttonA[1]*mb.buttonB[0])

	// dt == 0 -> linearly dependent, no consistent solution
	if dt == 0 {
		return 0
	}

	// Cramer's rule
	aa := przx*mb.buttonB[1] - przy*mb.buttonB[0]
	bb := przy*mb.buttonA[0] - przx*mb.buttonA[1]

	// Ensure integer values (partial presses are not possible)
	if aa%dt != 0 || bb%dt != 0 {
		return 0
	}

	// Required presses for A and B
	a := aa / dt
	b := bb / dt

	// If a or b is negative, there is no
	// solution (negative presses are not possible)
	if a >= 0 && b >= 0 {
		return 3*a + b // return cost
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res1 := int64(0)
	res2 := int64(0)

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
		mb.prizebase = [2]int64{10000000000000, 10000000000000}
		res2 += play(mb)

		if !el {
			break
		}
	}

	fmt.Println("res1:", res1)
	fmt.Println("res2:", res2)
}
