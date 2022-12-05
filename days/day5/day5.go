package day5

import (
	"adventofcode/m/v2/util"
	"fmt"
	"regexp"
)

func Day5(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Top Crates: %s\n", solve(inputFile))
	} else {
		fmt.Printf("Not implmenented.")
	}
}

func solve(inputFile string) string {

	ls := util.LineScanner(inputFile)
	
	stacks := []*Stack{}
	var numCrate int
	line, _ := util.Read(ls)
	nextLine, ok := util.Read(ls)

	firstLine := true
	for ok {
		if firstLine {
			stacks = append(stacks, &Stack{[]rune{}})
		}

		crateLabels := []rune(line)
		// 4 runes per create [ d ] _, we process 1, 5, 9, ...
		numCrate = 0
		for i := 1; i < len(crateLabels); i += 4 {
			if crateLabels[i] != ' ' {
				stacks[numCrate].push(crateLabels[i])
			}
			numCrate++
		}

		firstLine = false
		line = nextLine
		nextLine, ok = util.Read(ls)
		if nextLine == "" {
			break	
		}
	}

	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for ok {
		res := re1.FindStringSubmatch(line)
		toMove := strconv

		line, ok = util.Read(ls)
	}
}

type Stack struct { 
	crates []rune 
}

func (s *Stack) push(r rune) {
	s.crates = append(s.crates, r)
}

func (s *Stack) pop() rune {
	r := s.pop()
	s.crates = s.crates[:len(s-crates)-1]
	return r
}

func (s *Stack) head() rune {
	return s.crates[len(s.crates)-1]
}
