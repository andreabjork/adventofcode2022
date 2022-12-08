package day8

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
)

func Day8(inputFile string, part int) {
	f := MapForest(inputFile)
	if part == 0 {
		fmt.Printf("Visible: %d\n", f.countVisible())
	} else {
		fmt.Printf("Best view: %d\n", f.maxView())
	}
}

func MapForest(inputFile string) *Forest {
	ls := util.LineScanner(inputFile)
	forest := &Forest{map[int]map[int]*Tree{}}
	i := 0
	line, ok := util.Read(ls)
	for ok {
		forest.trees[i] = map[int]*Tree{}
		for j, h := range line {
			height, _ := strconv.Atoi(string(h))
			forest.trees[i][j] = &Tree{i, j, height, nil, nil, nil, nil}
		}
		i++
		line, ok = util.Read(ls)
	}
	return forest
}

type Forest struct {
	trees map[int]map[int]*Tree
}

type Tree struct {
	i int
	j int
	height int
	// The closest neighbour trees providing cover
	lCover *Tree
	rCover *Tree
	tCover *Tree
	bCover *Tree
}

func (f *Forest) maxView() int {
	bestView := 0
	for i := 0; i < len(f.trees); i++ {
		for j := 0; j < len(f.trees[0]); j++ {
			l,r,t,b := f.view(i,j)
			view := l*r*t*b
			if view > bestView {
				bestView = view
			}
		}
	}
	return bestView
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
	t, _ := f.get(i,j)
	lcov, rcov, tcov, bcov := f.cover(t)
	return lcov == nil || rcov == nil || tcov == nil || bcov == nil
}

func (f *Forest) get(i int, j int) (*Tree, bool) {
	if i >= 0 && i < len(f.trees) && j >= 0 && j < len(f.trees[i]) {
		return f.trees[i][j], true
	}
	// Tree out of bounds
	return nil, false
}

func (f *Forest) view(i int, j int) (int,int,int,int) {
	lv, rv, tv, bv := j, len(f.trees[i]) - j - 1, i, len(f.trees) - i - 1
	l,r,t,b := f.cover(f.trees[i][j])
	if l != nil {
		lv = j-l.j
	}
	if r != nil {
		rv = r.j-j
	}
	if t != nil {
		tv = i-t.i
	}
	if b != nil {
		bv = b.i-i
	}

	return lv,rv,tv,bv
}

func (f *Forest) cover(t *Tree) (*Tree, *Tree, *Tree, *Tree) {
	lc,rc,tc,bc := t.lCover, t.rCover, t.tCover, t.bCover

	// Find left cover
	if t.lCover == nil {
		for j := t.j-1; j >= 0; j-- {
			tt, found := f.get(t.i, j)
			if found && tt.height >= t.height {
				lc = tt
				break
			}
			if found && tt.lCover != nil && tt.lCover.height >= t.height {
				lc, _, _, _ = f.cover(tt)
				break
			}
		}
		t.lCover = lc
	}

	// Find right cover
	if t.rCover == nil {
		for j := t.j+1; j <= len(f.trees[t.i]); j++ {
			tt, found := f.get(t.i, j)
			if found && tt.height >= t.height {
				rc = tt
				break
			}
			if found && tt.rCover != nil && tt.rCover.height >= t.height {
				_, rc, _, _ = f.cover(tt)
				break
			}
		}
		t.rCover = rc
	}

	// Find top cover
	if t.tCover == nil {
		for i := t.i-1; i >= 0; i-- {
			tt, found := f.get(i, t.j)
			if found && tt.height >= t.height {
				tc = tt
				break
			}
			if found && tt.tCover != nil && tt.tCover.height >= t.height {
				_, _, tc, _ = f.cover(tt)
				break
			}
		}
		t.tCover = tc
	}

	// Find bottom cover
	if t.bCover == nil {
		for i := t.i+1; i <= len(f.trees); i++ {
			tt, found := f.get(i, t.j)
			if found && tt.height >= t.height {
				bc = tt
				break
			}
			if found && tt.bCover != nil && tt.bCover.height >= t.height {
				_, _, _, bc = f.cover(tt)
				break
			}
		}
		t.bCover = bc
	}

	return t.lCover, t.rCover, t.tCover, t.bCover
}