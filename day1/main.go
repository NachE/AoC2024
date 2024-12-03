// Copyright J.A. Nache. MIT license
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var aa sort.IntSlice
	var bb sort.IntSlice

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		ls := strings.Split(l, "   ")
		_an, err := strconv.Atoi(ls[0])
		if err != nil {
			panic(err)
		}
		aa = append(aa, _an)
		_bn, err := strconv.Atoi(ls[1])
		if err != nil {
			panic(err)
		}
		bb = append(bb, _bn)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sort.Sort(aa)
	sort.Sort(bb)

	r := 0
	r2 := 0
	for idx, av := range aa {
		fmt.Println(av, bb[idx])
		r = r + int(math.Abs(float64(av-bb[idx])))

		// part 2
		c := 0
		for _, bv := range bb {
			if bv == av {
				c++
			}
		}
		r2 = r2 + (av * c)
	}

	fmt.Println("res1:", r)
	fmt.Println("res2:", r2)
}
