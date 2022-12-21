package day16

import (
	"adventofcode/m/v2/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day16(inputFile string, part int) {
	if part == 0 {
		solve(inputFile)
	} else {
		fmt.Println("Not implmenented.")
	}
}

func solve(inputFile string) int {
	flow := parse(inputFile)
	// Compute distances from every node to every other node
	for i, _ := range flow.path {
		flow.dijkstra(i)
	}

	max := 0
	//max := flow.compute()
	//for v := 0; v < len(flow.valves); v++ {
	//	for w := v+1; w < len(flow.valves); w++ {
	//		alt := flow
	//		alt.path = flow.alternate(v, w)
	//		if f := alt.compute(); f > max {
	//			flow = alt
	//			max = f
	//		}
	//	}
	//}

	return max
}


type Flow struct {
	valves  map[string]*Valve // Ordered in the initial order
	path 	[]*Valve // Always a valid path
	active  []*Valve
	dist    map[*Valve]map[*Valve]int
	time    int
	flow    int
}

func (f *Flow) alternate(v, w int) []*Valve {
	path := []*Valve{}
	for i := 0; i < len(f.path); i++ {
		if i == v {
			path = append(path, f.path[w])
		} else if i == w {
			path = append(path, f.path[v])
		} else {
			path = append(path, f.path[i])
		}
	}
	return path
}

func (f *Flow) compute() int {
	f.time = 0
	for i := 0; i < len(f.path); i++ {
		f.open(f.path[i-1], f.path[i])
	}
	f.timePasses(30-f.time)

	return f.flow
}

func (f *Flow) open(from *Valve, to *Valve) {
	t := f.dist[from][to]
	f.timePasses(t+1)
	f.active = append(f.active, to)
}

func (f *Flow) timePasses(seconds int) {
	for _, ov := range f.active {
		f.flow += seconds*ov.rate
	}
}

type Valve struct {
	name      string
	rate      int
	tunnels   []*Valve // map[v] = x, where x is the length of path to v if one exists.
}

func parse(inputFile string) *Flow {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)
	re := regexp.MustCompile(`Valve ([A-Z]{2}) has flow rate=(\d+); tunnel[s]* lead[s]* to valve[s]* (.*)$`)

	f := &Flow {
		valves: map[string]*Valve{},
		path:   []*Valve{},
		active: []*Valve{},
		dist:   map[*Valve]map[*Valve]int{},
		time:   0,
		flow:   0,
	}

	ts := map[string][]string{}
	for ok {
		parts := re.FindStringSubmatch(line)
		fmt.Println(parts)
		valve := parts[1]
		rate, _ := strconv.Atoi(parts[2])
		tunnels := strings.Split(parts[3], ", ")
		v := &Valve{
			name:    valve,
			rate:    rate,
			tunnels: []*Valve{},
		}
		ts[valve] = tunnels
		f.valves[valve] = v
		f.path = append(f.path, v)

		line, ok = util.Read(ls)
	}

	for _, valve := range f.path {
		for _, t := range ts[valve.name] {
			valve.tunnels = append(valve.tunnels, f.valves[t])
		}
	}

	return f
}

// ==================
// Network & Dijkstra
// ==================


func (f *Flow) dijkstra(start int) {
	unvisited := append(f.path[:start], f.path[start+1:]...)
	// Iterate until all nodes have been visited
	i := 0
	for len(unvisited) > 0 {
		// Check all neighbours of unvisited node
		for k := 0; k < len(unvisited[i].tunnels); k++ {
			// If node is reachable, visit it
			if dist, ok := f.dist[f.path[start]][unvisited[i].tunnels[k]]; ok {
				f.dist[f.path[start]][unvisited[i]] = dist+1
				f.updateNeighbours(f.path[start], unvisited[i])
				unvisited = append(unvisited[0:i], unvisited[i+1:]...)
				i--
			}
		}
		i++
	}
}

func (f *Flow) updateNeighbours(start, v *Valve) {
	for k := 0; k < len(v.tunnels); k++ {
		if f.dist[start][v]+1 < f.dist[start][v.tunnels[k]] {
			f.dist[start][v.tunnels[k]] = f.dist[start][v]+1
			f.updateNeighbours(start, v.tunnels[k])
		}
	}
}
