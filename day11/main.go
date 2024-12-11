// Copyright J.A. Nache. MIT license
package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	if dat[len(dat)-1] == ' ' {
		dat = dat[:len(dat)-1]
	}
	dat = bytes.TrimSuffix(dat, []byte("\n"))
	_s := bytes.Split(dat, []byte{' '})
	stones_a := []int{}
	for _, st := range _s {
		ii, err := strconv.Atoi(string(st))
		if err != nil {
			panic(err)
		}
		stones_a = append(stones_a, ii)
	}

	for i := 0; i < 25; i++ {
		stones_b := []int{}
		for _, stone := range stones_a {
			if stone == 0 {
				stones_b = append(stones_b, 1)
				// strconv to prevent err on large numbers
			} else if len(strconv.Itoa(stone))%2 == 0 {
				nn := strconv.Itoa(stone)
				_a, err := strconv.Atoi(nn[:len(nn)/2])
				if err != nil {
					panic(err)
				}
				_b, err := strconv.Atoi(nn[len(nn)/2:])
				if err != nil {
					panic(err)
				}
				stones_b = append(stones_b, _a)
				stones_b = append(stones_b, _b)
			} else {
				stones_b = append(stones_b, stone*2024)
			}
		}
		stones_a = stones_b
	}

	fmt.Println("res1:", len(stones_a))
}
