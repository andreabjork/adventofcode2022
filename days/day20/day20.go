package day20

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
)

func Day20(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Grove coordinate sum: %d\n", solve(inputFile, 1, 1))
	} else {
		fmt.Printf("Grove coordinate sum: %d\n", solve(inputFile, 10, 811589153))
	}
}

func solve(inputFile string, times int, key int) int {
	ls := util.LineScanner(inputFile)

	var e, zero *Ele
	line, ok := util.Read(ls)
	// Grab the first
	num, _ := strconv.Atoi(line)
	first := &Ele{num*key, nil, nil}
	arr := []*Ele{first}

	prev := first
	line, ok = util.Read(ls)
	for ok {
		num, _ := strconv.Atoi(line)
		e = &Ele{num*key, nil, prev}

		prev.next = e
		prev = e
		if num == 0 {
			zero = e
		}
		arr = append(arr, e)
		line, ok = util.Read(ls)
	}

	// Connect the circle
	first.prev = e
	e.next = first

	arr = mix(arr, times)
	afterOT := zero.stepForward(1000)
	afterTT := afterOT.stepForward(1000)
	afterTHT := afterTT.stepForward(1000)
	return afterOT.val + afterTT.val + afterTHT.val
}


func mix(arr []*Ele, times int) []*Ele {
	for t := 0; t < times; t++ {
		for i := 0; i < len(arr); i++ {
			mixElement(arr[i], len(arr)-1)
		}
	}
	return arr
}

func mixElement(ele *Ele, N int) {
	if ele.val == 0 {
		return
	}
	d := ele.prev
	ele.remove()
	if ele.val > 0 {
		d = d.stepForward(ele.val%N)
	} else {
		d = d.stepBackward(ele.val%N)
	}

	ele.addAfter(d)
}

// ===========
// LINKED LIST
// ===========
type Ele struct {
	val   int
	next  *Ele
	prev  *Ele
}

func (e *Ele) remove() {
	e.prev.next = e.next
	e.next.prev = e.prev
}

func (e *Ele) addAfter(d *Ele) {
	e.prev = d
	e.next = d.next
	d.next.prev = e
	d.next = e
}

func (ele *Ele) stepForward(n int) *Ele {
	d := ele
	for i := 0; i < n; i++ {
		d = d.next
	}
	return d
}

func (ele *Ele) stepBackward(negN int) *Ele {
	d := ele
	for i := 0; i < util.Abs(negN); i++ {
		d = d.prev
	}
	return d
}
