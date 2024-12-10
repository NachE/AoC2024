// Copyright J.A. Nache. MIT license
package common

import (
	"fmt"
	"strconv"
)

func Byteint(b byte) int {
	r, err := strconv.Atoi(fmt.Sprintf("%v", string(b)))
	if err != nil {
		panic(err)
	}
	return r
}
