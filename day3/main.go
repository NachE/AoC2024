// Copyright J.A. Nache. MIT license
package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	regex := regexp.MustCompile(`(mul\((\d+?),(\d+?)\))|(do\(\))|(don\'t\(\))`)
	matches := regex.FindAllStringSubmatch(string((dat)), -1)

	sum := 0
	sump2 := 0
	b := true
	for _, m := range matches {
		switch m[0] {
		case "do()":
			b = true
			continue
		case "don't()":
			b = false
			continue
		}
		x, err := strconv.Atoi(m[2])
		if err != nil {
			fmt.Println(m)
			panic(err)
		}
		y, err := strconv.Atoi(m[3])
		if err != nil {
			panic(err)
		}
		_r := x * y
		sum += _r
		if b {
			sump2 += _r
		}
	}
	fmt.Println("part1 res:", sum)
	fmt.Println("part2 res:", sump2)
}
