// Copyright J.A. Nache. MIT license
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getmiddle(update []string) int {
	middle := update[((len(update)+1)/2)-1]
	mn, err := strconv.Atoi(middle)
	if err != nil {
		panic(err)
	}
	return mn
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rules := make(map[string]bool)

	// load rules
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lin := scanner.Text()
		if lin == "" {
			break
		}
		rules[lin[0:2]+lin[3:]] = true
	}

	res1 := 0
	res2 := 0
	// parse updates, part1
scanloop:
	for scanner.Scan() {
		lin := scanner.Text()
		update := strings.Split(lin, ",")
		for u := 0; u < len(update)-1; u++ {
			if _, exists := rules[update[u]+update[u+1]]; !exists {
				// not valid for part 1, proceed to part 2
				// sort the update for part 2
				sort.SliceStable(update, func(i, j int) bool {
					_, exists := rules[update[i]+update[j]]
					return exists
				})
				// get the middle for part 2
				mn := getmiddle(update)
				res2 += mn
				continue scanloop
			}
		}
		// get the middle for part 1
		mn := getmiddle(update)
		res1 += mn
	}

	fmt.Println("res1: ", res1)
	fmt.Println("res2: ", res2)
}
