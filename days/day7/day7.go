package day7

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
)

func Day7(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Total size: %d\n", solveA(inputFile))
	} else {
		fmt.Printf("Smallest large size: %d\n", solveB(inputFile))
	}
}



func solveA(inputFile string) int {
	fs := &FileSystem{  70000000, 0, root, []*Dir{root}}
	fs.examine(inputFile)
}

func solveB(inputFile string) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls) // skip $ cd / and just start from root
	line, ok = util.Read(ls)

	// Construct file system from root
	root := &Dir{"/", nil, -1, []int{}, map[string]*Dir{}}
	fs := &FileSystem{  70000000, 0, root, []*Dir{root}}
	dir := root
	for ok {
		dir = fs.interpret(line, dir)
		line, ok = util.Read(ls)
	}

	root.size()
	fs.used = root.bytes
	fmt.Println("Total: ", fs.total)
	fmt.Println("Used: ", fs.used)
	fmt.Println("Free: ", fs.total - fs.used)
	toRemove  := 30000000 - (fs.total - fs.used)
	fmt.Printf("We need to clean %d", toRemove)
	return fs.findMinOfMax(toRemove)
}

func solve(inputFile string, handler func() int) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls) // skip $ cd / and just start from root
	line, ok = util.Read(ls)

	// Construct file system from root
	root := &Dir{"/", nil, -1, []int{}, map[string]*Dir{}}
	fs := &FileSystem{70000000, 0, root, []*Dir{root}}
	dir := root
	for ok {
		dir = fs.interpret(line, dir)
		line, ok = util.Read(ls)
	}

	root.size()
	return fs.countSmall()
}

type FileSystem struct {
	total   int
	used 	int
	root 	*Dir
	dirs	[]*Dir
}

func (fs *FileSystem) countSmall() int {
	sum := 0
	for _, dir := range fs.dirs {
		if dir.bytes <= 100000 {
			sum += dir.bytes
		}
	}
	return sum
}

func (fs *FileSystem) findMinOfMax(max int) int {
	min := 70000000
	for _, dir := range fs.dirs {
		fmt.Println("Comparing ", dir.bytes, min)
		if dir.bytes >= max && dir.bytes < min {
			min = dir.bytes
		}
	}
	return min
}

// Line contains either an operation,
// or the result from an operation
func (fs *FileSystem) interpret(line string, dir *Dir) *Dir {
	fmt.Println(line)
	parts := strings.Split(line, " ")
	if len(parts) == 3 && parts[0] == "$" {
		if parts[1] == "ls" {
			return dir
		} else if parts[1] == "cd" && parts[2] == ".." {
			// Going back out, sum up the
			return dir.parent
		} else if parts[1] == "cd" {
			return dir.subdirs[parts[2]]
		}
	} else if len(parts) == 2 && parts[0] == "dir" {
		newDir := &Dir{parts[1], dir, -1, []int{}, map[string]*Dir{}}
		fs.dirs = append(fs.dirs, newDir)
		dir.subdirs[parts[1]] = newDir
	} else {
		size, _ := strconv.Atoi(parts[0])
		dir.files = append(dir.files, size)
	}

	return dir
}

// ===========
// DIRECTORIES
// ===========
type Dir struct {
	name        string
	parent		*Dir
	bytes 		int
	files  		[]int
	subdirs 	map[string]*Dir
}

func (d *Dir) size() int {
	total := 0
	for _, dir := range d.subdirs {
		total += dir.size()
	}
	d.bytes = total + sum(d.files)

	fmt.Printf("Size for %s: %d \n", d.name, d.bytes)
	return d.bytes
}

func sum(nums []int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}