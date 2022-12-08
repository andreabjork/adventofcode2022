package day8

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
)

func Day8(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Visible: %d\n", Solve(inputFile))
	} else {
		fmt.Println("Not implmenented.")
	}
}

func Solve(inputFile string) int {
	ls := util.LineScanner(inputFile)
	forest := &Forest{map[int]map[int]*Tree{}}

	i := 0
	line, ok := util.Read(ls)
	for ok {
		forest.trees[i] = map[int]*Tree{}
		for j, h := range line {
			height, _ := strconv.Atoi(string(h))
			l,r,t,b := -2,-2,-2,-2
			if i == 0 {
				t = -1
			}
			if j == 0 {
				l = -1
			} else if j == len(line)-1 {
				r = -1
			}
			forest.trees[i][j] = &Tree{i, j, height, l, r, t, b}
		}
		i++
		line, ok = util.Read(ls)
	}

	// Set all bottom covers to 0 in last line
	i--
	for j := 0; j < len(forest.trees[i]); j++ {
		forest.trees[i][j].bCover = -1
	}

	vis := forest.countVisible()
	forest.print()
	return vis
}

type Forest struct {
	trees map[int]map[int]*Tree
}

func (f *Forest) print() {
	for i := 0; i < len(f.trees); i++ {
		for j := 0; j < len(f.trees[i]); j++ {
			if f.visible(i,j) {
				fmt.Printf("V(%d) ", f.get(i,j).height)
			} else {
				fmt.Printf("H(%d) ", f.get(i,j).height)
			}
		}
		fmt.Printf("\n")
	}
}

type Tree struct {
	i int
	j int
	height  int
	lCover  int
	rCover int
	tCover int
	bCover int
}

func (f *Forest) countVisible() int {
	vis := 0
	for i := 0; i < len(f.trees); i++ {
		for j := 0; j < len(f.trees[0]); j++ {
			if f.visible(i, j) {
				vis++
			}
		}
	}
	return vis
}
func (f *Forest) visible(i int, j int) bool {
	t := f.get(i,j)
	lcov := f.leftCover(t)
	rcov := f.rightCover(t)
	tcov := f.topCover(t)
	bcov := f.bottomCover(t)

	return t.height > lcov || t.height > rcov || t.height > bcov || t.height > tcov
}

func (f *Forest) get(i int, j int) *Tree {
	if i >= 0 && i < len(f.trees) && j >= 0 && j < len(f.trees[i]) {
		return f.trees[i][j]
	}
	// Tree out of bounds, 0 height, 0 cover
	return &Tree{i, j, -1, 0, 0, 0, 0}
}

func (f *Forest) leftCover(t *Tree) int {
	if t.lCover == -2 {
		l := f.get(t.i, t.j-1)
		t.lCover = max(f.leftCover(l), l.height)
	}

	return t.lCover
}

func (f *Forest) rightCover(t *Tree) int {
	if t.rCover == -2 {
		l := f.get(t.i, t.j+1)
		t.rCover = max(f.rightCover(l), l.height)
	}

	return t.rCover
}

func (f *Forest) topCover(t *Tree) int {
	if t.tCover == -2 {
		l := f.get(t.i-1, t.j)
		t.tCover = max(f.topCover(l), l.height)
	}

	return t.tCover
}

func (f *Forest) bottomCover(t *Tree) int {
	if t.bCover == -2 {
		l := f.get(t.i+1, t.j)
		t.bCover = max(f.bottomCover(l), l.height)
	}

	return t.bCover
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}