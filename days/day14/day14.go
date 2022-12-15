package day14

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
)

func Day14(inputFile string, part int) {
	switch part {
	case 0:
		fmt.Printf("Grains of sand: %d\n", SimulateA(inputFile))
	case 1:
		fmt.Printf("Grains of sand: %d\n", SimulateB(inputFile))
	case 2:
		VisualizeA(inputFile)
	case 3:
		VisualizeB(inputFile)
	}
}

func SimulateA(inputFile string) int {
	c := MakeCave(inputFile, false)
	p := &Path{[]Step{}, 500, 0}
	// initial path
	OOB, _ := c.fall(p)
	i := 1 
	for !OOB  {
		p, OOB = c.next(p)
		if !OOB {
			i++
		}
	}
	c.trace(p).print()
	return i
}

func SimulateB(inputFile string) int {
	c := MakeCave(inputFile, true)
	p := &Path{[]Step{}, 500, 0}
	// initial path
	c.fall(p)
	i := 1 
	for !(p.endX == 500 && p.endY == 0)  {
		p, _ = c.next(p)
		i++
	}
	c.trace(p).print()
	return i
}

// ====
// CAVE
// ====
type Structure int64
const (
	Air   Structure = 0
	Rock            = 1
	Sand			= 2
	Trail           = 3 // Exists only for visualization purposes
)

type Cave struct {
	scan 				[][]Structure
	minX, maxX, maxY	int
}

const (
	SandColor = "\033[1;33m%s\033[0m"
	RockColor = "\033[1;31m%s\033[0m"
	PathColor = "\033[0;36m%s\033[0m"
)
func (c *Cave) print() {
	for y := 0; y <= c.maxY; y++ {
		for x := c.minX; x <= c.maxX; x++ {
			switch _, s := c.get(x,y); s {
			case Air:
				fmt.Printf(" ")
			case Rock:
				fmt.Printf(RockColor, "█")
			case Sand:
				fmt.Printf(SandColor, "*")
			case Trail:
				fmt.Printf(PathColor, "~")
			}
		}
		fmt.Println("")
	}
}

func (c *Cave) trace(p *Path) *Cave {
	ss := make([][]Structure, c.maxX-c.minX+1)
	for i := 0; i < len(ss); i++ {
		ss[i] = make([]Structure, c.maxY+1)
		for j := 0; j < len(ss[i]); j++ {
			ss[i][j] = c.scan[i][j]
		}
	}
	cc := &Cave{ss, c.minX, c.maxX, c.maxY}
	x := 500
	y := 0
	for i := 0; i < len(p.steps); i++ {
		switch p.steps[i] {
		case Left:
			x--
			y++
		case Right:
			x++
			y++
		case Down:
			y++
		}
		if x >= cc.minX && x <= cc.maxX && y >= 0 && y <= cc.maxY {
			cc.set(x, y, Trail)
		}
	}
	return cc
}

// Gets the next logical path q from a previous path p
func (c *Cave) next(p *Path) (*Path, bool) {
	steps := p.steps
	endX := p.endX
	endY := p.endY
	q := &Path{steps, endX, endY}
	// Alter the last step from previous path: If the last grain fell...
	// ... to the left and stopped, we can only fall to the right
	// ... to the right and stopped, we stop diagonally above
	// ... directly downwards, we fall to the left of it
	switch p.steps[len(p.steps)-1] {
	case Left:
		q.steps = p.steps[:len(p.steps)-1]
		q.endX++
		q.endY--
	case Right:
		q.steps = p.steps[:len(p.steps)-1]
		q.endX--
		q.endY--
	case Down:
		q.steps = p.steps[:len(p.steps)-1]
		q.endY--
	}

	// Keep falling downwards after that
	OOB, _ := c.fall(q)
	return q, OOB
}

func (c *Cave) fall(p *Path) (bool, bool) {
	blocked := true
	// Fall down until blocked or Out Of Bounds
	OOB, fallingDown := c.canFall(p, Down)
	for fallingDown && !OOB {
		blocked = false
		p.steps = append(p.steps, Down)
		p.endY++
		OOB, fallingDown = c.canFall(p, Down)
	}
	// Fall left or right otherwise
	OOB, canLeft := c.canFall(p, Left)
	OOB, canRight := c.canFall(p, Right)
	if canLeft {
		blocked = false
		p.steps = append(p.steps, Left)
		p.endX--
		p.endY++
	} else if canRight {
		blocked = false
		p.steps = append(p.steps, Right)
		p.endX++
		p.endY++
	}

	// Keep falling if not OOB or blocked
	if blocked {
		c.set(p.endX, p.endY, Sand)
		return OOB, blocked
	} else if OOB {
		return OOB, blocked
	} else {
		return c.fall(p)
	}
}

func (c *Cave) canFall(p *Path, dir Step) (bool, bool) {
	var OOB bool
	var unit Structure
	switch dir {
	case Left:
		OOB, unit = c.get(p.endX-1, p.endY+1)
	case Right:
		OOB, unit = c.get(p.endX+1, p.endY+1)
	case Down:
		OOB, unit = c.get(p.endX, p.endY+1)
	}
	return OOB, unit != Rock && unit != Sand
}

func (c *Cave) get(x, y int) (bool, Structure) {
	if x > c.maxX || x < c.minX || y > c.maxY {
		return true, Air
	} else {
		return false, c.scan[x-c.minX][y]
	}
}

func (c *Cave) set(x, y int, s Structure) {
	c.scan[x-c.minX][y] = s
}

// ====
// PATH
// ====
type Step int64
const (
	Left  Step = 0
	Right      = 1
	Down       = 2
)

// Falling path to (x,y)
// All paths start at 500,x
type Path struct {
	steps   	[]Step
	endX, endY  int
}

// =============
// PARSING INPUT
// =============
const MAX_INT = int(^uint(0) >> 1)
func MakeCave(inputFile string, infinite bool) *Cave {
	ls := util.LineScanner(inputFile)
	scan := [][]*Point{}
	maxX := -MAX_INT
	minX := MAX_INT
	maxY := -MAX_INT
	line, ok := util.Read(ls)
	i := 0
	for ok {
		scan = append(scan, []*Point{})
		parts := strings.Split(line, " -> ")
		for _, part := range parts {
			p := strings.Split(part, ",")
			x, _ := strconv.Atoi(p[0])
			y, _ := strconv.Atoi(p[1])

			if x > maxX {
				maxX = x
			}
			if x < minX {
				minX = x
			}
			if y > maxY {
				maxY = y
			}
			scan[i] = append(scan[i], &Point{x,y})
		}
		i++
		line, ok = util.Read(ls)
	}

	if infinite {
		maxX = maxX + 150
		minX = minX - 150
		maxY = maxY+2
	}

	// Now make the cave
	caveMap := make([][]Structure, maxX-minX+1)
	for i := 0; i < len(caveMap); i++ {
		caveMap[i] = make([]Structure, maxY+1)
	}

	c := &Cave{caveMap, minX, maxX, maxY}
	for _, s := range scan {
		i = 0
		for i+1 < len(s) {
			p := s[i]
			q := s[i+1]
			if p.x == q.x {
				for y := p.y; y <= q.y; y++ {
					c.set(p.x, y, Rock)
				}
				for y := q.y; y <= p.y; y++ {
					c.set(p.x, y, Rock)
				}
			}
			if p.y == q.y {
				for x := p.x; x <= q.x; x++ {
					c.set(x, p.y, Rock)
				}
				for x := q.x; x <= p.x; x++ {
					c.set(x, p.y, Rock)
				}
			}
			i++
		}
	}

	if infinite {
		for i := minX; i <= maxX; i++ {
			c.set(i, maxY, Rock)
		}
	}

	return c
}

type Point struct {
	x, y 	int
}
