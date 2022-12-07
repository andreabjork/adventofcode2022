package day7

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
)

func Day7(inputFile string, part int) {
	fs := newFS()
	fs.examine(inputFile)
	if part == 0 {
		fmt.Printf("Total size: %d\n", fs.totalSmall())
	} else {
		fmt.Printf("Smallest large size: %d\n", fs.makeSpace())
	}
}

func newFS() *FileSystem {
	root := &Dir{nil, -1, []int{}, map[string]*Dir{}}
	return &FileSystem{  70000000, 0, root, root, []*Dir{root}}
}

// ==============================================================
// FILESYSTEM
// --------------------------------------------------------------
// examine: reads input cmds, executes them and interprets output
// cd: change pwd into target directory
// mkdir: create directory at pwd
// touch: create file
// totalSmall: sums all dirs with size <= 100000
// makeSpace: deletes a single directory to have total free space
//			  at least 30000000
type FileSystem struct {
	total   int
	used 	int
	root 	*Dir
	pwd     *Dir
	dirs	[]*Dir
}

func (fs *FileSystem) examine(inputFile string) {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls) // skip $ cd / and just start from root
	line, ok = util.Read(ls)
	for ok {
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "$":
			if parts[1] == "cd" {
				fs.cd(parts[2])
			}
		case "dir":
			fs.mkdir(parts[1])
		default:
			size, _ := strconv.Atoi(parts[0])
			fs.touch(size)

		}
		line, ok = util.Read(ls)
	}
	fs.root.size()
}

func (fs *FileSystem) cd(folder string) {
	if folder == ".." {
		fs.pwd = fs.pwd.parent
	} else {
		fs.pwd = fs.pwd.subdirs[folder]
	}
}

func (fs *FileSystem) touch(fileSize int) {
	fs.pwd.files = append(fs.pwd.files, fileSize)
}

func (fs *FileSystem) mkdir(name string) {
	dir := &Dir{fs.pwd, -1, []int{}, map[string]*Dir{}}
	fs.dirs = append(fs.dirs, dir)
	fs.pwd.subdirs[name] = dir
}

func (fs *FileSystem) totalSmall() int {
	sum := 0
	for _, dir := range fs.dirs {
		if dir.bytes <= 100000 {
			sum += dir.bytes
		}
	}
	return sum
}

func (fs *FileSystem) makeSpace() int {
	fs.used = fs.root.bytes
	toRemove  := 30000000 - (fs.total - fs.used)
	min := fs.total
	for _, dir := range fs.dirs {
		if dir.bytes >= toRemove && dir.bytes < min {
			min = dir.bytes
		}
	}
	return min
}

// ===========
// DIRECTORIES
// ===========
type Dir struct {
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
	return d.bytes
}

func sum(nums []int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}