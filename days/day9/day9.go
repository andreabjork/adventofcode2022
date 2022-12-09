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

type Knot struct {
	x 			int
	y 			int
	path    	*Coord
	lookup    	map[int]map[int]bool
	visited 	int
}

func simulate(inputFile string, numKnots int) int {
	knots := []*Knot{}
	for i := 0; i < numKnots; i++ {
		knots = append(knots,
			&Knot{0,0, &Coord{0,0, nil},map[int]map[int]bool{0: {0: true}}, 1})
	}

	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)
	for ok {
		eles := strings.Split(line, " ")
		dir := []rune(eles[0])[0]
		N, _ := strconv.Atoi(eles[1])
		for i := 0; i < N; i++ {
			knots[0].move(dir)
			for j := 1; j < len(knots); j++ {
				knots[j].follow(knots[j-1])
			}
		}
		fmt.Println("")

		line, ok = util.Read(ls)
	}

	return knots[len(knots)-1].visited
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

	next := &Coord{k.x, k.y, k.path}
	if !k.lookup[k.x][k.y] {
		if k.lookup[k.x] == nil {
			k.lookup[k.x] = map[int]bool{}
		}
		k.lookup[k.x][k.y] = true
		k.visited++
	}
	k.path = next
	fmt.Printf("-> (%d,%d) ", k.x, k.y)
}

type Coord struct {
	x  			int
	y 			int
	from 		*Coord
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
