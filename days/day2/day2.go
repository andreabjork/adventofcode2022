package day2

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
)

func Day2(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Points: %d\n", RPS1(inputFile))
	} else {
		fmt.Printf("Points: %d\n", RPS2(inputFile))
	}
}

func RPS1(inputFile string) int {
	// Strategy guide specifies your and your opponents' play
	var points = map[string]int{
		// A,B,C = r,p,s (opponent); X,Y,Z = r,p,s (me)
		"X": 1,
		"Y": 2,
		"Z": 3,
		"AX": 3, // rock, rock = tie
		"AY": 6, // rock, paper = loss
		"AZ": 0, // rock, scissors = win
		"BX": 0, // paper, rock = win
		"BY": 3,
		"BZ": 6,
		"CX": 6, // scissors, rock = loss
		"CY": 0,
		"CZ": 3,
	}

	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	myPoints := 0
	for ok {
		p := strings.Split(line, " ")
		outcome := p[0]+p[1]
		myPoints += points[p[1]] + points[outcome]

		line, ok = util.Read(ls)
	}

	return myPoints
}

func RPS2(inputFile string) int {
	// Strategy guide specifies X,Y,Z = lose, draw, win
	var points = map[string]map[string]int {
		// I want to lose
		"X": {"A": 3+0, "B": 1+0, "C": 2+0},
		// I want to draw
		"Y": {"A": 1+3, "B": 2+3, "C": 3+3},
		// I want to win
		"Z": {"A": 2+6, "B": 3+6, "C": 1+6},
	}

	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	myPoints := 0
	for ok {
		p := strings.Split(line, " ")
		myPoints += points[p[1]][p[0]]

		line, ok = util.Read(ls)
	}

	return myPoints
}