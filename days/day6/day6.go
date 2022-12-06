package day6

import (
	"fmt"
	"adventofcode/m/v2/util"
)

func Day6(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Unique at: %d\n", solA(inputFile))
	} else {
		fmt.Println("Not implmenented.")
	}
}

func solA(inputFile string) int {
	rs := util.RuneScanner(inputFile)
	
	bufferSize := 4 
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
			return n
		}
		r, ok = util.Read(rs)
		n++
	}

	return 0
}

type asRune(str string) {
	return []rune(str)[0]
}

type Buffer struct {
	runes 		 []rune 
	iter		 	 int	  
	size 	  	 int	  
}

func (b *Buffer) shift(r rune) {
	b.runes[b.iter] = r
	b.iter++
	b.iter = b.iter%b.size
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
