package day6

import (
	"adventofcode/m/v2/util"
	"fmt"
)

func Day6(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Unique at: %d\n", sol(inputFile, 4))
	} else {
		fmt.Printf("Unique at: %d\n", sol(inputFile, 14))
	}
}

func sol(inputFile string, bufferSize int) int {
	rs := util.RuneScanner(inputFile)

	n := 0
	r, ok := util.Read(rs)

	init := make([]rune, bufferSize)
	for i := 0; i < bufferSize && ok; i++ {
		init[i] = asRune(r)
		r, ok = util.Read(rs)
		n++
	}

	buffer := &Buffer{init, 0, bufferSize}
	for ok {
		buffer.shift(asRune(r))
		if buffer.unique() {
			return n + 1
		}
		r, ok = util.Read(rs)
		n++
	}

	return 0
}

func asRune(str string) rune {
	return []rune(str)[0]
}

type Buffer struct {
	runes []rune
	iter  int
	size  int
}

func (b *Buffer) shift(r rune) {
	b.runes[b.iter] = r
	b.iter++
	b.iter = b.iter % b.size
}

func (b *Buffer) unique() bool {
	lookup := map[rune]bool{}
	for i := 0; i < len(b.runes); i++ {
		if lookup[b.runes[i]] {
			return false
		}

		lookup[b.runes[i]] = true
	}

	return true
}
