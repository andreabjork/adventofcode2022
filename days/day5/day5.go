package day5

import (
	"adventofcode/m/v2/util"
	"fmt"
	"regexp"
	"strconv"
)

func Day5(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Top Crates: %s\n", solve(inputFile, false))
	} else {
		fmt.Printf("Top Crates: %s\n", solve(inputFile, true))
	}
}

func solve(inputFile string, canMoveMany bool) string {
	ls := util.LineScanner(inputFile)
	// Read 2 lines at a time
	line, _ := util.Read(ls)
	nextLine, ok := util.Read(ls)

	// Process initial stacks
	stacks := []*Stack{}
	numCrate := 0
	for ok {
		crateLabels := []rune(line)
		// 4 runes per crate: '[ d ] _', we process 1, 5, 9, ...
		for i := 1; i < len(crateLabels); i += 4 {
			if len(stacks) < numCrate+1  {
				stacks = append(stacks, &Stack{[]rune{}})
			}
			if crateLabels[i] != ' ' {
				stacks[numCrate].crates = append([]rune{crateLabels[i]}, stacks[numCrate].crates...)
			}
			numCrate++
		}
		numCrate = 0
		line = nextLine
		nextLine, ok = util.Read(ls)
		if nextLine == "" {
			break	
		}
	}

	PrintStacks(stacks)

	line, ok = util.Read(ls)
	// Process movements
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for ok {
		res := re.FindStringSubmatch(line)
		toMove, _ := strconv.Atoi(res[1])
		fromStack, _ := strconv.Atoi(res[2])
		toStack, _ := strconv.Atoi(res[3])

		if canMoveMany {
			stacks[toStack-1].push(stacks[fromStack-1].pop(toMove))	
		} else {
			for i := 0;  i < toMove; i++ {
				stacks[toStack-1].push(stacks[fromStack-1].pop(1))
			}
		}
		

		//PrintStacks(stacks)
		line, ok = util.Read(ls)
	}

	tops := ""
	for i := 0; i < len(stacks); i++ {
		tops += string(stacks[i].peek())
	}

	return tops
}

func PrintStacks(stacks []*Stack)  {
	maxLength := 0
	for _, stack := range stacks {
		if len(stack.crates) > maxLength {
			maxLength = len(stack.crates)
		}
	}

	crateLevel := maxLength-1
	for crateLevel >= 0 {
		for i := 0; i < len(stacks); i++ {
			if len(stacks[i].crates) > crateLevel {
				fmt.Printf("[%s] ", string(stacks[i].crates[crateLevel]))
			} else {
				fmt.Printf("    ")
			}
		}
		fmt.Printf("\n")
		crateLevel--
	}
	
}
type Stack struct { 
	crates []rune 
}

func (s *Stack) push(r []rune) {
	s.crates = append(s.crates, r...)
}

func (s *Stack) pop(n int) []rune {
	popped := s.crates[len(s.crates)-n:]
	s.crates = s.crates[:len(s.crates)-n]
	return popped
}

func (s *Stack) peek() rune {
	return s.crates[len(s.crates)-1]
}
