package day12

import (
	"adventofcode/m/v2/util"
	"fmt"
)

func Day12(inputFile string, part int) {
	fixedStart, bestStart := solve(inputFile)
	if part == 0 {
		fmt.Printf("Steps to E: %d\n", fixedStart.shortestPath)
	} else {
		fmt.Printf("Steps to best S: %d\n", bestStart.shortestPath)
	}
}

func solve(inputFile string) (*Node, *Node) {
	d := &DTree{map[int]map[int]*Node{}, nil}
	unvisited, S, E := parse(inputFile)

	// Set starting node at E and find shortest path to every other node (including s)
	d.spt[E.x] = map[int]*Node{}
	d.spt[E.x][E.y] = E
	d.bestStart = E
	d.dijkstra(unvisited)
	return S, d.bestStart
}

const MAX_INT = int(^uint(0) >> 1)

// ===========================
// Dijkstra Shortest-Path-Tree
// ===========================
type DTree struct {
	spt 		map[int]map[int]*Node
	bestStart 	*Node
}

func (d *DTree) dijkstra(unvisited []*Node) {
	// Iterate until all nodes have been visited
	i := 0
	foundSome := false
	for len(unvisited) > 0 {
		if reachable, fromNode := d.findBestNeighbour(unvisited[i]); reachable {
			d.visit(unvisited[i], fromNode)
			d.updateBestStart(unvisited[i])
			d.updateNeighbours(unvisited[i])

			// remove node from unvisited
			unvisited = append(unvisited[0:i], unvisited[i+1:]...)
			i--
			foundSome = true
		}
		i++
		if i == len(unvisited) {
			i = 0
			if !foundSome {
				// some nodes might be unreachable
				break
			} else {
				foundSome = false
			}
		}
	}
}

func (d *DTree) visit(n *Node, fromNode *Node) {
	n.shortestPath = fromNode.shortestPath+1
	if d.spt[n.x] == nil {
		d.spt[n.x] = map[int]*Node{}
	}
	d.spt[n.x][n.y] = n
}

func (d *DTree) updateBestStart(n *Node) {
	if n.height < d.bestStart.height {
		d.bestStart = n
	} else if n.height == d.bestStart.height && n.shortestPath < d.bestStart.shortestPath {
		d.bestStart = n
	}
}

func (d *DTree) findBestNeighbour(n *Node) (bool, *Node) {
	minPath := MAX_INT
	var bestNeighbour *Node
	minPath, bestNeighbour = d.compareToNeighbour(n, n.x-1, n.y, minPath, bestNeighbour)
	minPath, bestNeighbour = d.compareToNeighbour(n, n.x+1, n.y, minPath, bestNeighbour)
	minPath, bestNeighbour = d.compareToNeighbour(n, n.x, n.y-1, minPath, bestNeighbour)
	minPath, bestNeighbour = d.compareToNeighbour(n, n.x, n.y+1, minPath, bestNeighbour)

	return bestNeighbour != nil, bestNeighbour
}

func (d *DTree) compareToNeighbour(n *Node, i, j, min int, best *Node) (int, *Node) {
	m, exists := d.spt[i][j]
	if exists && n.reachableFrom(m) {
		if m.shortestPath+1 < min {
			min = m.shortestPath+1
			best = m
		}
	}

	return min, best
}

func (d *DTree) updateNeighbours(n *Node) {
	// up, down, left right
	d.updateSingle(n,n.x-1,n.y)
	d.updateSingle(n,n.x+1,n.y)
	d.updateSingle(n,n.x,n.y-1)
	d.updateSingle(n,n.x,n.y+1)
}

func (d *DTree) updateSingle(here *Node, i, j int) {
	neighbour, exists := d.spt[i][j]
	if exists && neighbour.reachableFrom(here) && here.shortestPath + 1 < neighbour.shortestPath {
		neighbour.shortestPath = here.shortestPath+1
		d.updateBestStart(neighbour)
		d.updateNeighbours(neighbour)
	}
}

// ====
// Node
// ====
type Node struct {
	x				int
	y 				int
	height 			rune
	shortestPath	int
}

func (to *Node) reachableFrom(from *Node) bool {
	if int(to.height) < int(from.height) {
		return util.Abs(int(from.height)-int(to.height)) <= 1
	} else {
		return true
	}
}

// =============
// PARSING INPUT
// =============
func parse(inputFile string) ([]*Node, *Node, *Node) {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	var S, E *Node
	unvisited := []*Node{}
	i := 0
	M := 0
	for ok {
		runes := []rune(line)
		M = len(runes)
		for j := 0; j < M; j++ {
			r := runes[j]
			if r == 'S' {
				S = &Node {
					i,
					j,
					'a',
					MAX_INT,
				}
				unvisited = append(unvisited, S)
			} else if r == 'E' {
				E = &Node {
					i,
					j,
					'z',
					0,
				}
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

	return unvisited, S, E
}
