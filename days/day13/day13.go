package day13

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
)

func Day13(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Packets in right order: %d\n", solveA(inputFile))
	} else {
		fmt.Printf("Decoder Key: %d\n", solveB(inputFile))
	}
}

func solveA(inputFile string) int {
	ls := util.LineScanner(inputFile)
	lstr, _ := util.Read(ls)
	rstr, ok := util.Read(ls)
	sum := 0
	pairNum := 1
	for ok {
		left := &Packet{lstr, 0}
		right := &Packet{rstr, 0}
		ordered := left.leq(right)
		if ordered {
			sum += pairNum
		}
		pairNum += 1
		_, ok = util.Read(ls)
		lstr, _ = util.Read(ls)
		rstr, _ = util.Read(ls)
	}
	return sum
}

func solveB(inputFile string) int {
	ls := util.LineScanner(inputFile)
	key := &Key{[]*Packet{}}
	str, ok := util.Read(ls)
	for ok {
		left := &Packet{str, 0}
		key.place(left)
		str, ok = util.Read(ls)
		if str == "" {
			str, ok = util.Read(ls)
		}
	}
	key.place(&Packet{"[[2]]", 0})
	key.place(&Packet{"[[6]]", 0})
	div2 := -1
	div6 := -1
	for i, packet := range key.packets {
		if packet.bits == "[[2]]" {
			div2 = i
			if div6 > -1 {
				break
			}
		} else if packet.bits == "[[6]]" {
			div6 = i
			if div2 > -1 {
				break
			}
		}
	}
	return (div2+1)*(div6+1)
}

type Key struct {
	// packets is always ordered
	packets	[]*Packet
}

func (k *Key) place(p *Packet) {
	if len(k.packets) == 0 {
		k.packets = []*Packet{p}
		return
	}
	idx := len(k.packets)
	for i := 0; i < len(k.packets); i++ {
		ordered := p.leq(k.packets[i])
		if ordered {
			idx = i
			break
		}
	}
	// Create new ordered packet list
	if idx == len(k.packets) {
		k.packets = append(k.packets[:idx], p)
	} else {
		tail := []*Packet{}
		for i := idx; i < len(k.packets); i++{
			tail = append(tail, k.packets[i])
		}
		k.packets = append(append(k.packets[:idx], p), tail...)
	}
}

type Packet struct {
	bits 	 string
	iter 	 int
}

func (left *Packet) leq(right *Packet) bool {
	immutableLeft := left.bits
	immutableRight := right.bits
	ordered := true
	leftHasNext := left.hasNext()
	for leftHasNext {
		// right ran out of items
		if !right.hasNext() {
			ordered = false
			break
		}
		lNext := left.next()
		rNext := right.next()
		if lNext != rNext {
			leftInt, lIsInt := strconv.Atoi(lNext)
			rightInt, rIsInt := strconv.Atoi(rNext)
			// 		1. both are int (compare)
			if lIsInt == nil && rIsInt == nil {
				if leftInt < rightInt {
					break
				} else {
					ordered = false
					break
				}
			// 		2. neither is int (] , [) (ran out of items)
			} else if lIsInt != nil && rIsInt != nil {
				if !right.hasNext() || rNext == "]" {
					ordered = false
				}
				break
			//      3. either is int -> modify x to [ x ]
			} else if rIsInt != nil {
				left.push("["+lNext+"]")
				right.push(rNext)
			} else if lIsInt != nil {
				right.push("["+rNext+"]")
				left.push(lNext)
			}
		}
		leftHasNext = left.hasNext()
	}

	// Return the packets to their original state because they've been modified:
	left.bits = immutableLeft
	left.iter = 0
	right.bits = immutableRight
	right.iter = 0
	return ordered
}

func (p *Packet) print() {
	fmt.Println(p.bits)
}

func (p *Packet) next() string {
	ret := string(p.bits[p.iter])
	p.iter++

	// Support numbers > 9
	if p.hasNext() {
		_, err := strconv.Atoi(ret)
		_, err2 := strconv.Atoi(string(p.bits[p.iter]))
		if err == nil && err2 == nil {
			ret += string(p.bits[p.iter])
			p.iter++
			if p.iter >= len(p.bits) {
				err = nil
			} else {
				_, err2 = strconv.Atoi(string(p.bits[p.iter]))
			}
		}
	}
	return ret
}

func (p *Packet) hasNext() bool {
	return len(p.bits) > p.iter
}

func (p *Packet) push(bits string) {
	p.bits = p.bits[:p.iter-1] + bits + p.bits[p.iter:]
	p.iter--
}
