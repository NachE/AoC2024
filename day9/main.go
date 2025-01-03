// Copyright J.A. Nache. MIT license
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/NachE/AoC2024/day9/common"
	"github.com/NachE/AoC2024/day9/part2"
)

type bread struct {
	file *os.File
	pos  int64
	bid  []int64 // blockid stack
}

func (b *bread) pop() int64 {
	if len(b.bid) == 0 {
		// go to next valid pos
		b.file.Seek(b.pos, 0)
		// read block size value
		bff := make([]byte, 1)
		_, err := b.file.Read(bff)
		if err != nil {
			panic(err)
		}
		bsize := common.Byteint(bff[0])

		// build block file id
		fileid := b.pos / 2
		for i := 0; i < bsize; i++ {
			// store block file id * block size
			b.bid = append(b.bid, fileid)
		}
		b.pos = b.pos - 2 // store next valid pos
	}

	// pop
	v := b.bid[len(b.bid)-1]
	b.bid = b.bid[:len(b.bid)-1]
	return v
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		panic(err)
	}

	lbp := fi.Size()
	// warn: asumed odd valid char size
	if lbp%2 == 0 {
		// The last character is 10 (carriage return),
		// and the next one represents the file block size.
		// Moving back 2 positions gives us the last
		// file block size.
		lbp = lbp - 2
	} else {
		// back 3 positions [blocksize <-, freespace, carriagereturn]
		lbp = lbp - 3
	}

	bfile := &bread{
		file: file,
		pos:  lbp, // the last file block
	}

	var fwpos int64 = 0
	var rearangedpos int64 = -1
	// start reading file forward
	bff := make([]byte, 2)
	var res1 int64 = 0
	for {
		if fwpos > bfile.pos {
			break
		}
		file.Seek(fwpos, 0) // bring back to forward pos
		ff, err := file.Read(bff)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		// fwpos is also block id
		bsize := common.Byteint(bff[0])
		for i := 0; i < bsize; i++ {
			rearangedpos++
			res1 += (fwpos / 2) * int64(rearangedpos)
		}

		// move blocks to current free space
		freespace := common.Byteint(bff[1])
		for i := 0; i < freespace; i++ {
			rearangedpos++
			bwid := bfile.pop()
			res1 += bwid * int64(rearangedpos)
		}

		fwpos += int64(ff) // update next forward pos
	}

	// remain files
	for _, bwid := range bfile.bid {
		rearangedpos++
		res1 += bwid * int64(rearangedpos)
	}

	// part 2

	// The strategy used for part 1 is too complex to be applied to
	// part 2. Therefore, a new strategy should be implemented for part 2.
	// This is a reminder for future implementations: an overly specific
	// solution might only work for a particular problem and could be
	// difficult to refactor.

	res2 := part2.Main(file)

	fmt.Println("res1: ", res1)
	fmt.Println("res2: ", res2)
}
