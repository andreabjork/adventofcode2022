package day1

import (
	"adventofcode/m/v2/util"
	"fmt"
	"sort"
	"strconv"
)

func Day1(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Top 1 kcal: %d\n", day1a(inputFile))
	} else {
		fmt.Printf("Top 3 kcal: %d\n", day1b(inputFile))
	}
}

func day1a(inputFile string) int {
	t := topThree(inputFile)
	return t[len(t)-1]
}

func day1b(inputFile string) int {
	t := topThree(inputFile)
	return t[0]+t[1]+t[2]
}

func topThree(inputFile string) []int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	thisElf := 0
	topElves := []int{0,0,0}
	for ok {
		if line == "" {
			topElves = append(topElves, thisElf)
			sort.Ints(topElves)
			topElves = topElves[1:]
			thisElf = 0
		}

		kcal, _ := strconv.Atoi(line)
		thisElf += kcal

		line, ok = util.Read(ls)
	}

	return topElves
}
