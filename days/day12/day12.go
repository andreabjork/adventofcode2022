package day12

import (
	"adventofcode/m/v2/util"
	"fmt"
)

func Day12(inputFile string, part int) {
	if part == 0 {
		bestSignal := solve(inputFile)
		fmt.Printf("Best signal is at %s, reachable in steps: %d\n", string(bestSignal.height), bestSignal.shortestPath)
	} else {
		fmt.Println("Not implmenented.")
	}
}



func solve(inputFile string) *Node {
	d := &DTree{map[int]map[int]*Node{}, nil, 0, 0}
	unvisited, start, finish, rows, cols := parse(inputFile)
	d.rows = rows
	d.cols = cols

	// Set starting node
	d.spt[start.x] = map[int]*Node{}
	d.spt[start.x][start.y] = start
	d.bestSignal = start

	d.dijkstra(unvisited)
	d.show()
	return finish
}

const MAX_INT = int(^uint(0) >> 1)

type Node struct {
	x				int
	y 				int
	height 			rune
	shortestPath	int
}

type DTree struct {
	spt 		map[int]map[int]*Node
	// Because E is not always reachable, track this:
	bestSignal 	*Node
	rows 		int
	cols 		int
}

func (d *DTree) show() {
	for i := 0; i < d.rows; i++ {
		for j := 0; j < d.cols; j++ {
			if node, exists := d.spt[i][j]; exists {
				fmt.Printf("%d ", node.shortestPath)
			} else {
				// unreachable nodes
				fmt.Printf(". ")
			}
		}
		fmt.Println("")
	}

	for i := 0; i < d.rows; i++ {
		for j := 0; j < d.cols; j++ {
			if node, exists := d.spt[i][j]; exists {
				fmt.Printf("%s", string(node.height))
			} else {
				// unreachable nodes
				fmt.Printf(".")
			}
		}
		fmt.Println("")
	}
}

func (d *DTree) dijkstra(unvisited []*Node) {
	// Iterate until all nodes have been visited
	i := 0
	foundSome := false
	for len(unvisited) > 0 {
		if reachable, fromNode := d.findBest(unvisited[i]); reachable {
			// Visit node
			unvisited[i].shortestPath = fromNode.shortestPath+1
			if d.spt[unvisited[i].x] == nil {
				d.spt[unvisited[i].x] = map[int]*Node{}
			}
			// Add node to visited
			d.spt[unvisited[i].x][unvisited[i].y] = unvisited[i]
			d.updateBestSignal(unvisited[i])

			d.updateNeighbours(unvisited[i])
			// remove this from unvisited
			unvisited = append(unvisited[0:i], unvisited[i+1:]...)
			i--
			foundSome = true
		}
		i++
		if i == len(unvisited) {
			i = 0
			if !foundSome {
				break
			} else {
				foundSome = false
			}
		}
	}
}

func reachable(from, to *Node) bool {
	if int(to.height) > int(from.height) {
		return util.Abs(int(from.height)-int(to.height)) <= 1
	} else {
		return true
	}
}

func (d *DTree) updateBestSignal(n *Node) {
	if n.height > d.bestSignal.height {
		d.bestSignal = n
	} else if n.height == d.bestSignal.height && n.shortestPath < d.bestSignal.shortestPath {
		d.bestSignal = n
	}
}
func (d *DTree) findBest(n *Node) (bool, *Node) {
	up, uExists := d.spt[n.x-1][n.y]
	down, dExists := d.spt[n.x+1][n.y]
	left, lExists := d.spt[n.x][n.y-1]
	right, rExists := d.spt[n.x][n.y+1]

	minPath := MAX_INT
	var bestNeighbour *Node
	if lExists && reachable(left, n) {
		if left.shortestPath+1 < minPath {
			minPath = left.shortestPath+1
			bestNeighbour = left
		}
	}

	if rExists && reachable(right, n) {
		if right.shortestPath+1 < minPath {
			minPath = right.shortestPath+1
			bestNeighbour = right
		}
	}

	if uExists && reachable(up, n) {
		if up.shortestPath+1 < minPath {
			minPath = up.shortestPath+1
			bestNeighbour = up
		}
	}

	if dExists && reachable(down, n) {
		if down.shortestPath+1 < minPath {
			minPath = down.shortestPath+1
			bestNeighbour = down
		}
	}

	return bestNeighbour != nil, bestNeighbour
}

func (d *DTree) updateSingle(here *Node, i, j int) {
	neighbour, exists := d.spt[i][j]
	if exists && reachable(here, neighbour) && here.shortestPath + 1 < neighbour.shortestPath {
		neighbour.shortestPath = here.shortestPath+1
		d.updateBestSignal(neighbour)
		d.updateNeighbours(neighbour)
	}
}

func (d *DTree) updateNeighbours(n *Node) {
	// up
	d.updateSingle(n,n.x-1,n.y)
	// down
	d.updateSingle(n,n.x+1,n.y)
	// left
	d.updateSingle(n,n.x,n.y-1)
	// right
	d.updateSingle(n,n.x,n.y+1)
}

// =============
// PARSING INPUT
// =============
func parse(inputFile string) ([]*Node, *Node, *Node, int, int) {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	var start, finish *Node
	unvisited := []*Node{}
	i := 0
	M := 0
	for ok {
		runes := []rune(line)
		M = len(runes)
		for j := 0; j < M; j++ {
			r := runes[j]
			if r == 'S' {
				start = &Node {
					i,
					j,
					'a',
					0,
				}
			} else if r == 'E' {
				finish = &Node {
					i,
					j,
					'z',
					MAX_INT,
				}
				unvisited = append(unvisited, finish)
			} else {
				unvisited = append(unvisited,
					&Node{
						i,
						j,
						r,
						MAX_INT,
					},
				)
			}
		}
		line, ok = util.Read(ls)
		i++
	}
	rows := i
	cols := M

	return unvisited, start, finish, rows, cols
}
