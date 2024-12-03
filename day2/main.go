// Copyright J.A. Nache. MIT license
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tolerate(ls []string) bool {
	for i := 0; i < len(ls); i++ {
		ls2 := make([]string, 0)
		ls2 = append(ls2, ls[:i]...)
		ls2 = append(ls2, ls[i+1:]...)
		if isSafe(ls2) {
			return true
		}
	}
	return false
}

func isSafe(ls []string) bool {
	ant, err := strconv.Atoi(ls[0])
	if err != nil {
		panic(err)
	}

	increase := false
	decrease := false
	for _, av := range ls[1:] {
		avn, err := strconv.Atoi(av)
		if err != nil {
			panic(err)
		}

		diff := ant - avn
		if diff == 0 {
			return false
		}

		if diff < 0 {
			increase = true
		}

		if diff > 0 {
			decrease = true
		}

		if diff > 3 || diff < -3 {
			return false
		}

		if increase && decrease {
			return false
		}

		ant = avn
	}

	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safe := 0
	safer2 := 0
	for scanner.Scan() {
		l := scanner.Text()
		ls := strings.Split(l, " ")

		if isSafe(ls) {
			safe++
			safer2++
		} else {
			if tolerate(ls) {
				safer2++
			}
		}
	}

	fmt.Println("res1:", safe)
	fmt.Println("res2:", safer2)
}
