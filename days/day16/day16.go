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
	fmt.Printf("Initial path: %+v\n", flow.path)
	for i, _ := range flow.path {
		fmt.Printf("%d: %s\n", i, flow.path[i].name)
		flow.dijkstra(i)
	}
	for v, val := range flow.dist {
		for w, d := range val {
			fmt.Printf("dist %s to %s: %d\n", v.name, w.name, d)
		}
	}

	//max := 0
	fmt.Println("Before: ", len(flow.path))
	flow.path = flow.initialize()
	max := flow.compute()
	fmt.Println("Initial flow is ", max)
	
	for midx := 1; midx < len(flow.path)-1; midx++ {
		for v := midx; v < len(flow.path); v++ {
			for w := v+1; w < len(flow.path); w++ {
				// TODO this needs to be more robust
				alt := flow
				alt.path = flow.alternate(v, w)
				if f := alt.compute(); f > max {
					flow = alt
					max = f
				}
				fmt.Println("-------")
			}
		}
	}

	fmt.Println("BEST FLOW IS THIS ONE")
	for i := 0; i < len(flow.path); i++ {
		fmt.Printf("%s -> ", flow.path[i].name)
	}
	fmt.Println(max)
	return max
}

type Flow struct {
	valves map[string]*Valve // Ordered in the initial order
	path   []*Valve          // Always a valid path
	active []*Valve
	dist   map[*Valve]map[*Valve]int
	time   int
	flow   int
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

func (f *Flow) initialize() []*Valve {
	path := []*Valve{f.path[0]}
	for i := 1; i < len(f.path); i++ {
		if f.path[i].rate > 0 {
			path = append(path, f.path[i])
		}
	}
	return path
}

func (f *Flow) compute() int {
	f.flow = 0
	f.time = 0
	f.active = []*Valve{}
	if f.path[0].rate > 0 {
		f.open(f.valves["AA"], f.path[0])
	}
	for i := 0; i < len(f.path)-1; i++ {
		if f.path[i+1].rate > 0 {
			f.open(f.path[i], f.path[i+1])
		}
	}
	fmt.Printf("All open, now passing %d time\n", 31-f.time)
	n := 30-f.time
	for i := 0; i < n; i++ {
		f.timePasses(1)
	}
	fmt.Printf("Computed flow: t=%d, f=%d\n", f.time, f.flow)
	return f.flow
}

func (f *Flow) open(from *Valve, to *Valve) {
	t, ok := f.dist[from][to]
	if !ok{
		fmt.Println("FATAL")
	} else {
		fmt.Printf("Heading %s -> %s ( %d minutes )\n", from.name, to.name, t+1)
	}
	for i := 0; i < t+1; i++ {
		f.timePasses(1)
	}
	f.active = append(f.active, to)
	fmt.Printf("Opened %s; t=%d, f: %d\n", to.name, f.time, f.flow)
}

func (f *Flow) timePasses(seconds int) {
	f.time += seconds
	fmt.Printf("t=%d | Valves ", f.time)
	press := 0
	for _, ov := range f.active {
		fmt.Printf("%s, ", ov.name)
		f.flow += seconds * ov.rate
		press += seconds * ov.rate
	}
	fmt.Printf("are open, releasing %d pressure\n", press)
}

type Valve struct {
	name    string
	rate    int
	tunnels []*Valve // map[v] = x, where x is the length of path to v if one exists.
}

func parse(inputFile string) *Flow {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)
	re := regexp.MustCompile(`Valve ([A-Z]{2}) has flow rate=(\d+); tunnel[s]* lead[s]* to valve[s]* (.*)$`)

	f := &Flow{
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

	fmt.Printf("AA: %d\n", f.valves["AA"].rate)
	for _, t := range f.valves["AA"].tunnels {
		fmt.Printf("-> %s\n", t.name)
	}
	return f
}

// ==================
// Network & Dijkstra
// ==================
func (f *Flow) dijkstra(start int) {
	if f.dist[f.path[start]] == nil {
		f.dist[f.path[start]] = map[*Valve]int{}
		for _, t := range f.path[start].tunnels {
			f.dist[f.path[start]][t] = 1
		}
		f.dist[f.path[start]][f.path[start]] = 0
	}

	unvisited := copy(f.path[start+1:])
	unvisited = append(unvisited, f.path[:start]...)
	// Iterate until all nodes have been visited
	i := 0
	for len(unvisited) > 0 {
		// Check all neighbours of unvisited node
		for k := 0; k < len(unvisited[i].tunnels); k++ {
			// If node is reachable, visit it
			if dist, ok := f.dist[f.path[start]][unvisited[i].tunnels[k]]; ok {
				if initDist, ok := f.dist[f.path[start]][unvisited[i]]; ok && initDist > dist+1 {
					// distance was already set in the initialization step
						f.dist[f.path[start]][unvisited[i]] = dist + 1
				} else if !ok {
						f.dist[f.path[start]][unvisited[i]] = dist + 1
				}
				f.updateNeighbours(f.path[start], unvisited[i])
				// remove element
				arr := unvisited[:i]
				unvisited = append(arr, unvisited[i+1:]...)
				if len(unvisited) == 0 {
					break
				} else {
					i = i % len(unvisited)
				}
			}
		}
		i++
		if i >= len(unvisited) && len(unvisited) > 0 {
			i = i % len(unvisited)
		}
	}
}

func (f *Flow) updateNeighbours(start, v *Valve) {
	for k := 0; k < len(v.tunnels); k++ {
		if f.dist[start][v]+1 < f.dist[start][v.tunnels[k]] {
			f.dist[start][v.tunnels[k]] = f.dist[start][v] + 1
			f.updateNeighbours(start, v.tunnels[k])
		}
	}
}

func copy(path []*Valve) []*Valve {
	npath := []*Valve{}
	for i := 0; i < len(path); i++ {
		npath = append(npath, path[i])
	}
	return npath
}