package day17

//import (
//	"adventofcode/m/v2/util"
//	"fmt"
//)
//
//func Day17(inputFile string, part int) {
//	if part == 0 {
//		fmt.Printf("Tower height: %d\n", solve(inputFile, 2022))
//	} else {
//		fmt.Println("Not implmenented.")
//	}
//}
//
//func solve(inputFile string, NUM_ROCKS int) int {
//	hmap := []int{0,0,0,0,0,0,0}
//	gas := parse(inputFile)
//	order := []func() *Rock{
//		bar,
//		plus,
//		revL,
//		pole,
//		box,
//	}
//	g := &Game{hmap, gas, 0, order}
//	for i := 0; i < NUM_ROCKS; i++ {
//		g.drop(i%len(g.order))
//		g.print()
//	}
//
//	return g.max()
//}
//
//type Game struct {
//	hmap	[]int
//	gas 	[]rune
//	t 	    int
//	order   []func() *Rock
//}
//
//
//func (g *Game) max() int {
//	max := 0
//	for i := 0; i < len(g.hmap); i++ {
//		if g.hmap[i] > max {
//			max = g.hmap[i]
//		}
//	}
//	return max
//}
//
//func (g *Game) drop(r int) {
//	rock := g.order[r]()
//	rock.setTo(g.max()+4)
//	//fmt.Println("--- DROP ---")
//	rock.print()
//	i := 0
//	for true {
//		fmt.Println(string(g.gas[g.t]))
//		g.push(rock, g.gas[g.t])
//		g.t = (g.t+1)%len(g.gas)
//		if g.stopped(rock) {
//			rock.print()
//			break
//		} else {
//			rock.dropOne()
//			rock.print()
//		}
//		i++
//	}
//	fmt.Println("stack height: ", g.max())
//}
//
//func (g *Game) push(r *Rock, dir rune) {
//	switch dir {
//	case '<':
//		// If there is space to the left, push left
//		if g.spaceLeftOf(r) {
//			r.hmap = append(r.hmap[1:], -1)
//			for i := 0; i < len(r.pixels); i++ {
//				r.pixels[i] = append(r.pixels[i][1:], 0)
//			}
//		}
//	case '>':
//		if g.spaceRightOf(r) {
//			r.hmap = append([]int{-1}, r.hmap[:len(r.hmap)-1]...)
//			for i := 0; i < len(r.pixels); i++ {
//				r.pixels[i] = append([]int{0}, r.pixels[i][:len(r.pixels[i])-1]...)
//			}
//		}
//	}
//}
//
//func (g *Game) spaceLeftOf(r *Rock) bool {
//	blocked := false
//	if r.hmap[0] != -1 {
//		// blocked by wall
//		blocked = true
//	}
//
//	for j := 1; j < len(r.hmap); j++ {
//		// blocked by rocks
//		if r.hmap[j] != -1 {
//			if g.hmap[j-1] >= r.hmap[j] {
//				blocked = true
//			}
//			break
//		}
//	}
//	return !blocked
//}
//
//func (g *Game) spaceRightOf(r *Rock) bool {
//	blocked := false
//	if r.hmap[len(r.hmap)-1] != -1 {
//		// blocked by wall
//		blocked = true
//	}
//
//	for j := len(r.hmap)-2; j >= 0; j-- {
//		// blocked by rocks
//		if r.hmap[j] != -1 {
//			if g.hmap[j+1] >= r.hmap[j] {
//				blocked = true
//			}
//			break
//		}
//	}
//	return !blocked
//}
//
//func (g *Game) stopped(r *Rock) bool {
//	fmt.Printf("Compare for: %+v\n", g.hmap)
//	fmt.Printf("Compare to: %+v\n", r.hmap)
//	clashHeight := -1
//	for i := 0; i < len(g.hmap); i++ {
//		if g.hmap[i]+1 == r.hmap[i] {
//			clashHeight = g.hmap[i]+1
//		}
//	}
//	if clashHeight >= 0 {
//		tops := r.tops(clashHeight-(clashHeight-r.min()))
//		for i := 0; i < len(g.hmap); i++ {
//			g.hmap[i] = util.Max(g.hmap[i], tops[i])
//		}
//	}
//	return clashHeight >= 0
//}
//
//// ====
//// ROCK
//// ====
//type Rock struct {
//	pixels [][]int
//	hmap   []int
//}
//
//
//func (r *Rock) dropOne() {
//	for i := 0; i < len(r.hmap); i++ {
//		if r.hmap[i] > 0 {
//			r.hmap[i]--
//		}
//	}
//}
//
//func (r *Rock) setTo(h int) {
//	hmap := []int{-1,-1,-1,-1,-1,-1,-1}
//	for i := len(r.pixels)-1; 	i >= 0; i-- {
//		for j := 0; j < len(r.pixels[i]); j++ {
//			if r.pixels[i][j] > 0 {
//				r.pixels[i][j] = h + (len(r.pixels)-1-i)
//			}
//			if r.pixels[i][j] > 0 && hmap[j] == -1 {
//				hmap[j] = r.pixels[i][j]
//			}
//		}
//	}
//	r.hmap = hmap
//}
//
//func (r *Rock) tops(h int) []int {
//	// Example for the + rock:
//	// clashIndex 0 ->  2,3,2
//	// clashIndex 1 -> 1,2,1
//	// clashIndex 2 -> 0,1,0
//	tops := []int{0,0,0,0,0,0,0}
//	for i := 0;	i < len(r.pixels); i++ {
//		for j := 0; j < len(r.pixels[i]); j++ {
//			if r.pixels[i][j] > 0 && tops[j] == 0 {
//				tops[j] = len(r.pixels) - i + h-1
//			}
//		}
//	}
//
//	return tops
//}
//
//const MAX_INT = int(^uint(0) >> 1)
//func (r *Rock) min() int {
//	min := MAX_INT
//	for i := 0; i < len(r.hmap); i++ {
//		if r.hmap[i] < min && r.hmap[i] != -1 {
//			min = r.hmap[i]
//		}
//	}
//	return min
//}
//
//func bar() *Rock {
//	return &Rock{[][]int{
//		{0, 0, 1, 1, 1, 1, 0},
//		},
//		[]int{},
//	}
//}
//func plus() *Rock {
//	return &Rock{[][]int{
//		{0, 0, 0, 1, 0, 0, 0},
//		{0, 0, 1, 1, 1, 0, 0},
//		{0, 0, 0, 1, 0, 0, 0},
//		},
//		[]int{},
//	}
//}
//func revL() *Rock {
//	return &Rock{[][]int{
//		{0, 0, 0, 0, 1, 0, 0},
//		{0, 0, 0, 0, 1, 0, 0},
//		{0, 0, 1, 1, 1, 0, 0},
//		},
//		[]int{},
//	}
//}
//
//func pole() *Rock {
//	return &Rock{[][]int{
//			{0, 0, 1, 0, 0, 0, 0},
//			{0, 0, 1, 0, 0, 0, 0},
//			{0, 0, 1, 0, 0, 0, 0},
//			{0, 0, 1, 0, 0, 0, 0},
//		},
//		[]int{},
//	}
//}
//
//func box() *Rock {
//	return &Rock{[][]int{
//		{0, 0, 1, 1, 0, 0, 0},
//		{0, 0, 1, 1, 0, 0, 0},
//		},
//		[]int{},
//	}
//}
//
//// PARSING
//func parse(inputFile string) []rune {
//	ls := util.LineScanner(inputFile)
//	line, _ := util.Read(ls)
//	return []rune(line)
//}
//
//// DEBUG
//func (r *Rock) print() {
//	for i := 0; i < len(r.hmap); i++ {
//		if r.hmap[i] == -1 {
//			fmt.Printf(".")
//		} else {
//			fmt.Printf("%d", r.hmap[i])
//		}
//	}
//	fmt.Println("")
//}
//
//func (g *Game) print() {
//	fmt.Println("%+v\n", g.hmap)
//}