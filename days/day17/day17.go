package day17

import (
	"adventofcode/m/v2/util"
	"fmt"
)

func Day17(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Tower height: %d\n", solve(inputFile, 2022))
	} else {
		fmt.Printf("Tower height: %d\n", solve(inputFile, 1000000000000))
	}
}

func solve(inputFile string, NUM_ROCKS int) int {
	// Empty board with 7 comlumns
	board := []map[int]int{{},{},{},{},{},{},{}}
	gas := parse(inputFile)
	order := []func() *Rock{bar, plus, revL, pole, box,}
	g := &Game{board, gas, order, 0, 0}
	intervals := []int{}
	prev := 0
	var rocksDropped int
	repeats := 0
	height := 0
	for rocksDropped = 0; rocksDropped < NUM_ROCKS; rocksDropped++ { // Drop NUM_ROCKS or until we reach equilibrium
		g.drop(rocksDropped%len(g.order))

		// State reaches equilibrium where every repeating x5 rocks will be dropped in the same order, same manner,
		// result in the same height added to the existing stack.
		if rocksDropped % 5 == 0 {
			intervals = append(intervals, g.max-prev)
			prev = g.max
			eq, reps, stack := checkForEquilibrium(intervals)
			if eq {
				repeats = reps
				height = stack
				break
			}
		}
	}
	eqHeight := 0
	if rocksDropped < NUM_ROCKS {
		// Quick drop by tallying up height of repeating stacks
		mult := (NUM_ROCKS-rocksDropped)/(repeats*5)
		quickDrop := (mult-1)*repeats*5
		eqHeight = (mult-1)*height
		// Drop the rest
		for i := rocksDropped+quickDrop+1; i < NUM_ROCKS; i++ {
			g.drop(i%len(g.order))
		}
	}

	return g.max+eqHeight
}

func checkForEquilibrium(heights []int) (bool, int, int) {
	// Start with window as the back item
	window := []int{heights[len(heights)-1]}
	for i := len(heights)-2; i > len(heights)/3 + len(window); i-- {
		if len(heights) >= 3*len(window) {
			// Compare window to the next 2 iterations
			firstStart := len(heights)-2*len(window)
			secondStart := len(heights)-3*len(window)
			eq := true
			for k := 0; k < len(window); k++ {
				if !(window[k] == heights[firstStart+k] && window[k] == heights[secondStart+k]) {
					eq = false
					break
				}
			}
			// If equilibrium was not reached with that, extend window and repeat
			if !eq || len(window) < 5 {
				window = append([]int{heights[i]}, window...)
			} else {
				return true, len(window), util.Sum(window)
			}
		}
	}

	return false, 0, 0
}


type Game struct {
	board   []map[int]int
	gas 	[]rune
	order   []func() *Rock
	t 	    int
	max 	int
}

func (g *Game) drop(rType int) {
	r := g.order[rType]()
	r.height(g.max+4)
	ok := true
	//g.print(r)
	for ok {
		_ = g.shift(r, g.gas[g.t])
		g.t = (g.t+1)%len(g.gas)
		ok = g.shift(r, 'd')
		if !ok {
			//g.print(r)
			g.add(r)
			//fmt.Println("Rock stopped")
			//g.printBoard()
		} else {
			//g.print(r)
		}
	}
}

// ====
// ROCK
// ====
type Rock struct {
	pixels [][]int
}

func (g *Game) add(r *Rock) {
	for i := 0; i < len(r.pixels); i++ {
		for j := 0; j < len(r.pixels[i]); j++ {
			if r.pixels[i][j] > 0 {
				g.board[j][r.pixels[i][j]]	= 1
			}
			g.max = util.Max(r.pixels[i][j],g.max)
		}
	}
}

func (r *Rock) height(h int) {
	for i := len(r.pixels)-1; 	i >= 0; i-- {
		for j := 0; j < len(r.pixels[i]); j++ {
			r.pixels[i][j] = r.pixels[i][j]*(h + (len(r.pixels)-1-i))
		}
	}
}

func (r *Rock) copy() *Rock {
	pixs := [][]int{}
	for i := 0; i < len(r.pixels); i++ {
		pixs = append(pixs, []int{})
		for j := 0; j < len(r.pixels[i]); j++ {
			pixs[i] = append(pixs[i], r.pixels[i][j])
		}
	}
	return &Rock{pixs}
}

func (g *Game) shift(r *Rock, pivot rune) bool {
	pr := r.copy()
	switch pivot {
	case 'd':
		for i := 0; i < len(pr.pixels); i++ {
			for j := 0; j < len(pr.pixels[i]); j++ {
				if _, occupied := g.board[j][pr.pixels[i][j]-1]; occupied || pr.pixels[i][j]-1 == 0 {
					return false
				}
				if pr.pixels[i][j] > 0 {
					pr.pixels[i][j]--
				}
			}
		}
	case '<':
		// Check if already at edge
		for i := 0; i < len(pr.pixels); i++ {
			if pr.pixels[i][0] != 0 {
				return false
			}
		}
		// Else shift if there is space
		for i := 0; i < len(pr.pixels); i++ {
			for j := 0; j < len(pr.pixels[i])-1; j++ {
				if _, occupied := g.board[j][pr.pixels[i][j+1]]; occupied {
					return false
				}

				pr.pixels[i][j] = pr.pixels[i][j+1]
			}
			pr.pixels[i][len(pr.pixels[i])-1] = 0
		}
	case '>':
		// Check if already at edge
		for i := 0; i < len(pr.pixels); i++ {
			if pr.pixels[i][len(pr.pixels[0])-1] != 0 {
				return false
			}
		}
		// Else shift if there is space
		for i := 0; i < len(pr.pixels); i++ {
			for j := len(pr.pixels[i])-1; j > 0; j-- {
				if _, occupied := g.board[j][pr.pixels[i][j-1]]; occupied {
					return false
				}
				pr.pixels[i][j] = pr.pixels[i][j-1]
			}
			pr.pixels[i][0] = 0
		}
	}

	r.pixels = pr.pixels
	return true
}

func bar() *Rock {
	return &Rock{[][]int{
		{0, 0, 1, 1, 1, 1, 0},
		},
	}
}
func plus() *Rock {
	return &Rock{[][]int{
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		},
	}
}
func revL() *Rock {
	return &Rock{[][]int{
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0},
		},
	}
}

func pole() *Rock {
	return &Rock{[][]int{
			{0, 0, 1, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
		},
	}
}

func box() *Rock {
	return &Rock{[][]int{
		{0, 0, 1, 1, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0},
		},
	}
}

// PARSING
func parse(inputFile string) []rune {
	ls := util.LineScanner(inputFile)
	line, _ := util.Read(ls)
	return []rune(line)
}

// DEBUG
func (r *Rock) print() {
	for i := 0; i < len(r.pixels); i++ {
		for j := 0; j < len(r.pixels[i]); j++{
			fmt.Printf("%d ", r.pixels[i][j])
		}
		fmt.Println("")
	}
}
func (g *Game) print(r *Rock) {
	rock := []map[int]int{{},{},{},{},{},{},{},}
	for i := 0; i < len(r.pixels); i++ {
		for j := 0; j < len(r.pixels[i]); j++ {
			if r.pixels[i][j] > 0 {
				rock[j][r.pixels[i][j]] = 1
			}
		}
	}
	i := 0
	for h := g.max+4; h > 0; h-- {
		fmt.Printf("| ")
		for col := 0; col < 7; col ++ {
			if _, exists := g.board[col][h]; exists {
				fmt.Printf("# ")
			} else if _, exists = rock[col][h]; exists {
				fmt.Printf("@ ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Printf("|")
		i++
		fmt.Println()
	}
	fmt.Println("+ - - - - - - - +")
}

// DEBUG
func (g *Game) printBoard() {
	for h := g.max+3; h > 0; h--{
		fmt.Printf("| ")
		for col := 0; col < 7; col ++ {
			if _, exists := g.board[col][h]; exists {
				fmt.Printf("# ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Printf("|")
		fmt.Println()
	}
	fmt.Println("+ - - - - - - - +")
}
