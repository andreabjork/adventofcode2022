package day9

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
)

func Day9(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("# Visited: %d\n", simulate(inputFile, 2))
	} else {
		fmt.Printf("# Visited: %d\n", simulate(inputFile, 10))
	}
}

func simulate(inputFile string, numKnots int) int {
	knots := []*Knot{}
	for i := 0; i < numKnots; i++ {
		knots = append(knots,
			&Knot{0,0,map[int]map[int]bool{0: {0: true}}, 1})
	}

	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)
	for ok {
		eles := strings.Split(line, " ")
		dir := []rune(eles[0])[0]
		N, _ := strconv.Atoi(eles[1])
		// For each step, move the head and let the rest follow
		for i := 0; i < N; i++ {
			knots[0].move(dir)
			for j := 1; j < len(knots); j++ {
				knots[j].follow(knots[j-1])
			}
		}

		line, ok = util.Read(ls)
	}

	return knots[len(knots)-1].visited
}

// =====
// KNOTS
// =====
type Knot struct {
	x 			int
	y 			int
	lookup    	map[int]map[int]bool
	visited 	int
}

func (k *Knot) move(dir rune) {
	switch dir {
	case 'U':
		k.y++
	case 'D':
		k.y--
	case 'L':
		k.x--
	case 'R':
		k.x++
	}
}

func (k *Knot) follow(h *Knot) {
	dist := abs(h.x - k.x) + abs(h.y - k.y)
	closeX := abs(h.x - k.x) <= 1
	closeY := abs(h.y - k.y) <= 1

	// move diagonally
	if dist > 2 {
		if h.x > k.x {
			k.move('R')
		} else {
			k.move('L')
		}
		if h.y > k.y {
			k.move('U')
		} else {
			k.move('D')
		}
	} else if !closeX && h.x > k.x {
		k.move('R')
	} else if !closeX && h.x < k.x {
		k.move('L')
	} else if !closeY && h.y > k.y {
		k.move('U')
	} else if !closeY && h.y < k.y {
		k.move('D')
	}

	k.trackVisit()
}

func (k *Knot) trackVisit() {
	if !k.lookup[k.x][k.y] {
		if k.lookup[k.x] == nil {
			k.lookup[k.x] = map[int]bool{}
		}
		k.lookup[k.x][k.y] = true
		k.visited++
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
