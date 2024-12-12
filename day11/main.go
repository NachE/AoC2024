// Copyright J.A. Nache. MIT license
package main

import (
	"bytes"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

// For part 1, a loop over an updated stack was sufficient,
// but for part 2, some optimization is needed: the algorithm
// has been modified to be recursive and to process one of
// the initial stones at a time. This frees up RAM, which
// we'll use to cache numbers and their results according to
// the number of blinks. This, in turn, reduces CPU usage,
// ultimately resulting in an efficient solution.
//
// More efficiency strategies can be implemented, such as using
// threads to leverage all of the CPU, but with this approach,
// it's sufficient since on a Mac Pro M1 it solves in 0.2s.
var cache map[string]int64 = make(map[string]int64)

func blink(stone int64, times int) int64 {
	ckey := fmt.Sprintf("%v-%v", stone, times)
	if _, exists := cache[ckey]; exists {
		return cache[ckey]
	}

	zero := big.NewInt(0)
	one := big.NewInt(1)
	n24 := big.NewInt(2024)

	bstone := big.NewInt(stone)

	for i := 0; i < times; i++ {
		if bstone.Cmp(zero) == 0 {
			bstone = one
			// strconv to prevent err on large numbers
		} else if len(bstone.String())%2 == 0 {
			nn := bstone.String()
			_a, err := strconv.ParseInt(nn[:len(nn)/2], 10, 0)
			if err != nil {
				panic(err)
			}
			_ra := blink(_a, times-(i+1))
			_b, err := strconv.ParseInt(nn[len(nn)/2:], 10, 0)
			if err != nil {
				panic(err)
			}
			_rb := blink(_b, times-(i+1))
			cache[ckey] = _ra + _rb
			return _ra + _rb
		} else {
			bstone.Mul(bstone, n24)
		}
	}

	cache[ckey] = 1
	return 1
}

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
	stones_a := []int64{}
	for _, st := range _s {
		ii, err := strconv.ParseInt(string(st), 10, 0)
		if err != nil {
			panic(err)
		}
		stones_a = append(stones_a, ii)
	}

	res1 := int64(0)
	for _, st := range stones_a {
		res1 += blink(st, 25)
	}

	res2 := int64(0)
	for _, st := range stones_a {
		res2 += blink(st, 75)
	}

	fmt.Println("res1:", res1)
	fmt.Println("res2:", res2)
}
