package day15

import (
	"adventofcode/m/v2/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day15(inputFile string, part int) {
	if part == 0 {
		if strings.Contains(inputFile, "0.txt") {
			fmt.Printf("Known spots: %d\n", solveA(inputFile, 2000000))
		} else if strings.Contains(inputFile, "1.txt") {
			fmt.Printf("Known spots: %d\n", solveA(inputFile, 10))
		}
	} else {
		if strings.Contains(inputFile, "0.txt") {
			fmt.Printf("Tuning Frequency: %d\n", solveB(inputFile, 4000001))
		} else if strings.Contains(inputFile, "1.txt") {
			fmt.Printf("Tuning Frequency: %d\n", solveB(inputFile, 21))
		}
	}
}

func solveA(inputFile string, y int) int {
	n := parse(inputFile)
	knowns := n.mapStrip(y)
	return knowns
}

func solveB(inputFile string, max int) int {
	n := parse(inputFile)
	return n.locate(max, max).tf()
}

type Network struct {
	sensors []*Sensor
	minX 	int
	minY 	int
	width 	int
	height  int
}

type Sensor struct {
	loc 	*Point
	beacon 	*Point
	dist 	int
}

type Point struct {
	x,y		int
}

func (n *Network) locate(maxX, maxY int) *Point {
	for y := util.Max(0, n.minY); y < util.Min(n.minY+n.width, maxY); y++ {
		if p := n.locateInStrip(y, maxX); p != nil {
			return p
		}
	}
	return nil
}

func (n *Network) locateInStrip(y int, maxX int) *Point {
	found := false
	foundBeacon := false
	foundSensor := false

	for x := util.Max(0,n.minX); x < util.Min(maxX, n.minX+n.width); x++ {
		for _, s := range n.sensors {
			if s.beacon.x == x && s.beacon.y == y {
				foundBeacon = true
			} else if s.loc.x == x && s.loc.y == y {
				foundSensor = true
			} else if s.loc.d(&Point{x,y}) <= s.dist {
				found = true
				break
			}
		}

		if !found && !foundBeacon && !foundSensor {
			return &Point{x, y}
		}
		found = false
		foundBeacon = false
		foundSensor = false
	}
	return nil
}

func (n *Network) mapStrip(y int) int {
	found := false
	knowns := 0
	for x := n.minX; x < n.minX+n.width; x++ {
		for _, s := range n.sensors {
			if !(s.beacon.x == x && s.beacon.y == y) && !(s.loc.x == x && s.loc.y == y) && s.loc.d(&Point{x,y}) <= s.dist {
				found = true
				break
			}
		}
		if found {
			knowns++
		}
		found = false
	}
	return knowns
}

func (p *Point) tf() int {
	return p.x*4000000 + p.y
}
func (p *Point) d(q *Point) int {
	// s is at (x,y) and detects beacon (p,q)
	// (a,b) is in range if |x-a|+|y-b| <= |p-a|+|q-b|
	return util.Abs(p.x-q.x)+util.Abs(p.y-q.y)
}


const MAX_INT = int(^uint(0) >> 1)
func parse(inputFile string) *Network {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)
	re := regexp.MustCompile(`Sensor at x=([-]*[0-9]+), y=([-]*[0-9]+): closest beacon is at x=([-]*[0-9]+), y=([-]*[0-9]+)`)

	n := &Network{
		//area:    [][]Structure{},
		sensors: []*Sensor{},
		minX:    MAX_INT,
		minY:    MAX_INT,
		width:   0,
		height:  0,
	}

	maxX, maxY := -MAX_INT, -MAX_INT
	for ok {
		m := re.FindStringSubmatch(line)
		sX, _ := strconv.Atoi(m[1])
		sY, _ := strconv.Atoi(m[2])
		bX, _ := strconv.Atoi(m[3])
		bY, _ := strconv.Atoi(m[4])
		sensorLoc := &Point{sX, sY}
		beaconLoc := &Point{bX, bY}

		n.sensors = append(n.sensors,
			&Sensor{
			loc:    sensorLoc,
			beacon: beaconLoc,
			dist:   sensorLoc.d(beaconLoc),
		})

		n.minX = util.Min(bX, util.Min(sX, n.minX))
		n.minY = util.Min(bY, util.Min(sY, n.minY))
		maxX = util.Max(bX, util.Max(sX, maxX))
		maxY = util.Max(bY, util.Max(sY, maxY))

		line, ok = util.Read(ls)
	}
	n.width = maxX - n.minX+1
	n.height = maxY - n.minY+1

	return n
}