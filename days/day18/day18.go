package day18

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
)

func Day18(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Surface area: %d\n", solve(inputFile, false))
	} else {
		fmt.Printf("Surface area: %d\n", solve(inputFile, true))
	}
}

const MAX_INT = int(^uint(0) >> 1)
func solve(inputFile string, exteriorOnly bool) int {
	a := &Area{map[int]map[int]map[int]int{}, 0, nil, MAX_INT, -MAX_INT, MAX_INT, -MAX_INT, MAX_INT, -MAX_INT}
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)
	for ok {
		p := toPoint(line)
		a.add(p)
		line, ok = util.Read(ls)
	}

	if exteriorOnly {
		a.findTrappedAir()
		return a.surface - a.trapped.surface
	} else {
		return a.surface
	}
}

type Area struct {
	points                        map[int]map[int]map[int]int
	surface                       int
	trapped                       *Area
	minX,maxX,minY,maxY,minZ,maxZ int
}

func (a *Area) findTrappedAir() {
	for x := a.minX; x <= a.maxX; x++{
		for y := a.minY; y <= a.maxY; y++ {
			for z := a.minZ; z <= a.maxZ; z++ {
				expanse := &Area{map[int]map[int]map[int]int{}, 0, nil, MAX_INT, -MAX_INT, MAX_INT, -MAX_INT, MAX_INT, -MAX_INT}
				unbounded := a.expandFrom(&Point{x,y,z}, expanse)
				if !unbounded {
					a.addTrapped(expanse)
				}
			}
		}
	}
}

func (a *Area) addTrapped(b *Area) {
	if a.trapped == nil {
		a.trapped = &Area{map[int]map[int]map[int]int{}, 0, nil, MAX_INT, -MAX_INT, MAX_INT, -MAX_INT, MAX_INT, -MAX_INT}
	}
	for x, p := range b.points {
		for y, pp := range p {
			for z, _ := range pp {
				a.trapped.add(&Point{x,y,z})
			}
		}
	}

}

func (a *Area) expandFrom(p *Point, expanse *Area ) bool {
	// If we are expanding from somewhere outside the area, then we have escaped and the air is not trapped
	if _, visited := expanse.points[p.x][p.y][p.z]; visited {
		return false
	} else if p.x < a.minX || p.x > a.maxX || p.y < a.minY || p.y > a.maxY || p.z < a.minZ || p.z > a.maxZ {
		return true
	} else if _, filled := a.points[p.x][p.y][p.z]; filled {
		// block is filled by obsidian, can not expand from here
		return false
	} else {
		expanse.add(p)
	}

	return a.
		expandFrom(&Point{p.x-1, p.y, p.z}, expanse) || a.
		expandFrom(&Point{p.x+1, p.y, p.z}, expanse) || a.
		expandFrom(&Point{p.x, p.y-1, p.z}, expanse) || a.
		expandFrom(&Point{p.x, p.y+1, p.z}, expanse) || a.
		expandFrom(&Point{p.x, p.y, p.z-1}, expanse) || a.
		expandFrom(&Point{p.x, p.y, p.z+1}, expanse)
}

func (a *Area) add(p *Point) {
	a.maxX = util.Max(a.maxX, p.x)
	a.minX = util.Min(a.minX, p.x)
	a.maxY = util.Max(a.maxY, p.y)
	a.minY = util.Min(a.minY, p.y)
	a.maxZ = util.Max(a.maxZ, p.z)
	a.minZ = util.Min(a.minZ, p.z)

	if a.points[p.x] == nil {
		a.points[p.x] = map[int]map[int]int{}
		a.points[p.x][p.y] = map[int]int{}
		a.points[p.x][p.y][p.z] = 1
	} else if a.points[p.x][p.y] == nil {
		a.points[p.x][p.y] = map[int]int{}
		a.points[p.x][p.y][p.z] = 1
	} else if _, ok := a.points[p.x][p.y][p.z]; ok {
		return
	} else {
		a.points[p.x][p.y][p.z] = 1
	}

	// Track surface edges
	if _, exists := a.points[p.x-1][p.y][p.z]; exists { a.surface-- } else { a.surface++ }
	if _, exists := a.points[p.x+1][p.y][p.z]; exists { a.surface-- } else { a.surface++ }
	if _, exists := a.points[p.x][p.y-1][p.z]; exists { a.surface-- } else { a.surface++ }
	if _, exists := a.points[p.x][p.y+1][p.z]; exists { a.surface-- } else { a.surface++ }
	if _, exists := a.points[p.x][p.y][p.z-1]; exists { a.surface-- } else { a.surface++ }
	if _, exists := a.points[p.x][p.y][p.z+1]; exists { a.surface-- } else { a.surface++ }
}

type Point struct {
	x,y,z int
}

func toPoint(line string) *Point {
	parts := strings.Split(line, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])

	return &Point{x,y,z}
}