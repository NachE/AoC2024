#!/bin/bash

for DAY in {1..25}
do
	mkdir day$DAY
	cd day$DAY
	go mod init github.com/NachE/AoC2024/day$DAY
	cat << 'EOF' > main.go
// Copyright J.A. Nache. MIT license
package main
import "fmt"

func main() {
    fmt.Println("Hello, world.")
}

EOF
cd -
done
