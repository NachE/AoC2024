// Copyright J.A. Nache. MIT license
package part2

import (
	"io"
	"os"

	"github.com/NachE/AoC2024/day9/common"
)

type sector interface {
	Size() int64
}

type vfile struct {
	id    int64
	size  int64
	moved bool // to prevent move file twice
}

func (v *vfile) Size() int64 {
	return v.size
}

type empty struct {
	size int64
}

func (e *empty) Size() int64 {
	return e.size
}

type vchunk struct {
	vsectors []sector
}

// salloc return true if there is empty sector >= size
func (v *vchunk) salloc(size int64) bool {
	for _, vs := range v.vsectors {
		if em, ok := vs.(*empty); ok {
			if em.size >= size {
				return true
			}
		}
	}
	return false
}

// allocatefile allocates vf in empty sector with size >= vf.size
func (v *vchunk) allocatefile(vf *vfile) bool {
	for si, vs := range v.vsectors {
		if em, ok := vs.(*empty); ok {
			if em.size >= vf.size {
				em.size = em.size - vf.size
				v.vsectors = append(v.vsectors[:si+1], v.vsectors[si:]...)
				v.vsectors[si] = vf
			}
		}
	}
	return false
}

type vdisk struct {
	vchunks []*vchunk
}

// biasleft vove vch chunk to left if there is space left.
func (v *vdisk) biasleft(vch *vchunk) {
	for si, vs := range vch.vsectors { // for every file in vch
		vf, ok := vs.(*vfile)
		if !ok || vf.moved {
			continue
		}

		// look for available vsector in left chunks for vf file
		for i := 0; i < len(v.vchunks); i++ {
			if v.vchunks[i] == vch {
				// skip self chunk and next chunks
				return
			}

			// chunk without space
			if !v.vchunks[i].salloc(vf.size) {
				continue
			}

			v.vchunks[i].allocatefile(vf)
			vch.vsectors[si] = &empty{size: vf.size}
			return
		}
	}
}

func Main(file *os.File) int64 {
	file.Seek(0, 0)

	bf := make([]byte, 2)
	vdsk := &vdisk{}

	id := int64(0)
	for {
		_, err := file.Read(bf)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		vch := &vchunk{}
		vf := &vfile{id: id, size: int64(common.Byteint(bf[0]))}
		vch.vsectors = append(vch.vsectors, vf)
		// warn: assume that the last byte is '\n' on freespace pos
		if bf[1] == '\n' { // end of file
			bf[1] = '0'
		}
		freespace := int64(common.Byteint(bf[1]))
		em := &empty{size: freespace}
		vch.vsectors = append(vch.vsectors, em)
		vdsk.vchunks = append(vdsk.vchunks, vch)
		id++
	}

	// travel chunks from right to left
	for i := len(vdsk.vchunks) - 1; i > 0; i-- {
		// try to move every chunk to empty space on left
		vdsk.biasleft(vdsk.vchunks[i])
	}

	// calc puzle
	res2 := int64(0)
	pos := int64(-1)
	for _, vchk := range vdsk.vchunks {
		for _, vs := range vchk.vsectors {
			if vf, ok := vs.(*vfile); ok {
				for _i := int64(0); _i < vf.size; _i++ {
					pos++
					res2 += pos * vf.id
				}
			}
			if em, ok := vs.(*empty); ok {
				pos += em.size
			}
		}
	}

	return res2
}
