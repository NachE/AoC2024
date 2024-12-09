// Copyright J.A. Nache. MIT license
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type state struct {
	idx int
	val int
}

func validate(nums []int, target int, opers []string) bool {
	// initial stack state
	stck := []state{{
		idx: 1,
		val: nums[0],
	}}

	for len(stck) > 0 {
		// pop one
		stt := stck[len(stck)-1]
		stck = stck[:len(stck)-1]

		// validate result
		if stt.idx == len(nums) {
			if stt.val == target {
				return true
			}
			continue
		}

		// next num
		nn := nums[stt.idx]

		for _, oper := range opers {
			switch oper {
			case "+":
				stck = append(stck, state{
					idx: stt.idx + 1,
					val: stt.val + nn,
				})
			case "*":
				stck = append(stck, state{
					idx: stt.idx + 1,
					val: stt.val * nn,
				})
			case "||":
				vvs := fmt.Sprintf("%v%v", stt.val, nn)
				vv, err := strconv.Atoi(vvs)
				if err != nil {
					panic(err)
				}
				stck = append(stck, state{
					idx: stt.idx + 1,
					val: vv,
				})
			}
		}
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	res1 := 0
	res2 := 0
	for scanner.Scan() {
		lin := scanner.Bytes()
		s := bytes.Split(lin, []byte(": "))

		target, err := strconv.Atoi(string(s[0]))
		if err != nil {
			panic(err)
		}
		bnums := bytes.Split(s[1], []byte(" "))
		nums := []int(nil)
		for _, bnum := range bnums {
			num, err := strconv.Atoi(string(bnum))
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		if validate(nums, target, []string{"*", "+"}) {
			res1 += target
		}
		if validate(nums, target, []string{"*", "+", "||"}) {
			res2 += target
		}
	}

	fmt.Println("res1:", res1)
	fmt.Println("res2:", res2)
}
