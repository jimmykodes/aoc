package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	p := puzzle("assets/input.txt")
	// p1(p)
	p = consolidate(p)
	p2(p)
}

func consolidate(b []Block) []Block {
	var out []Block
	var idx int
	for idx < len(b) {
		out = append(out, b[idx])
		idx += b[idx].Size
	}
	return out
}

func p2(blocks []Block) {
	rp := len(blocks) - 1
	for {
		for !blocks[rp].IsFile {
			// jump empty blocks
			// we know the first block is a file, no need to ensure rp >= 0
			rp--
		}
		if rp == 0 {
			break
		}
		fileBlock := blocks[rp]
		lp := 0
		for {
			// find the first empty space large enough for the file
			if lp < len(blocks) && (blocks[lp].IsFile || blocks[lp].Size < fileBlock.Size) {
				lp++
			} else {
				break
			}
		}
		if lp > rp {
			// we moved to the right of the file we are trying to move, skip the file and move on
			rp--
			continue
		}
		// we have a valid hole to the left of the file
		gap := blocks[lp]
		remainingSize := gap.Size - fileBlock.Size
		blocks[lp] = fileBlock
		// TODO: compact this with the empties on either side?
		blocks[rp] = Block{Size: fileBlock.Size}
		if remainingSize > 0 {
			// we need to insert a new empty block to the right of the file we moved
			blocks = slices.Insert(blocks, lp+1, Block{Size: remainingSize})
		}
	}
	var i int
	var checksum int
	for _, block := range blocks {
		for range block.Size {
			if block.IsFile {
				checksum += block.File * i
			}
			i++
		}
	}
	fmt.Println(checksum)
}

func p1(blocks []Block) {
	var (
		lp = 0
		rp = len(blocks) - 1
	)
	for {
		for blocks[lp].IsFile {
			// advance lp to next empty
			lp++
		}
		for !blocks[rp].IsFile {
			rp--
		}
		if lp >= rp {
			break
		}
		// swap
		blocks[lp], blocks[rp] = blocks[rp], blocks[lp]
	}
	var checksum int
	for i, block := range blocks {
		if !block.IsFile {
			break
		}
		checksum += block.File * i
	}
	fmt.Println(checksum)
}

type Block struct {
	IsFile bool
	File   int
	Size   int
}

func puzzle(filename string) []Block {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	var out []Block
	charToInt := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	fileID := -1
	for i, chunk := range data {
		size := charToInt[chunk]
		isFile := i%2 == 0
		if isFile {
			fileID++
		}
		for range size {
			if isFile {
				out = append(out, Block{IsFile: isFile, Size: size, File: fileID})
			} else {
				out = append(out, Block{IsFile: isFile, Size: size})
			}
		}
	}
	return out
}
