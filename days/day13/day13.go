package day13

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
)

func Day13(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Packets in right order: %d\n", solve(inputFile))
	} else {
		fmt.Println("Not implmenented.")
	}
}

func solve(inputFile string) int {
	ls := util.LineScanner(inputFile)
	lstr, _ := util.Read(ls)
	rstr, ok := util.Read(ls)
	sum := 0
	pairNum := 1
	for ok {
		left := &Packet{[]rune(lstr), 0}
		right := &Packet{[]rune(rstr), 0}
		fmt.Printf("\nLEFT\n")
		for lb := 0; lb < len(left.bits); lb ++ {
			fmt.Printf("%s ", string(left.bits[lb]))
		}
		fmt.Printf("\nRIGHT\n")
		for rb := 0; rb < len(right.bits); rb ++ {
			fmt.Printf("%s ", string(right.bits[rb]))
		}
		ordered := true
		leftHasNext := left.hasNext()
		for leftHasNext {
			// right ran out of items
			if !right.hasNext() {
				fmt.Println("Right ran out of items first")
				ordered = false
				break
			}
			// 7 = 7, [ = [, ] = ]
			lNext := left.next()
			rNext := right.next()
			//fmt.Printf("Comparing: %s, %s\n", string(lNext), string(lNext))
			if lNext != rNext { // unequal runes -> 3 different possibilities
				if lNext == ']' {
					// left ran out first
					break
				} else if rNext == ']' {
					ordered = false
					break
				}
				leftInt, lIsInt := strconv.Atoi(string(lNext))
				rightInt, rIsInt := strconv.Atoi(string(rNext))
				if lIsInt == nil && rIsInt == nil {
					// if unequal ints, compare value
					if leftInt < rightInt {
						break
					} else {
						ordered = false
						break
					}
				} else if lIsInt != nil && rIsInt != nil {
					// neither value is an integer (, ] [)
					if left.hasNext() && !right.hasNext() {
						ordered = false
					}
					break
				} else if rIsInt != nil {
					// right is not int
					left.push('[', lNext, ']')
					right.push(rNext)

					fmt.Println("")
					fmt.Println("Packets after adjusting: ")
					left.print()
					right.print()
				} else if lIsInt != nil {
					// left is not int
					right.push('[', rNext, ']')
					left.push(lNext)
					fmt.Println("")
					fmt.Println("Packets after adjusting: ")
					left.print()
					right.print()
				}
			}
			leftHasNext = left.hasNext()
		}
		// left ran out of items first: ok

		fmt.Println("Ordered?", ordered)
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

type Packet struct {
	bits 	 []rune
	iter 	 int
}

func (p *Packet) print() {
	for lb := 0; lb < len(p.bits); lb ++ {
		fmt.Printf("%s ", string(p.bits[lb]))
	}
	fmt.Println("")
	fmt.Println("Iterator at: ", p.iter)
}

func (p *Packet) next() rune {
	p.iter++
	return p.bits[p.iter-1]
}

func (p *Packet) hasNext() bool {
	return len(p.bits) > p.iter
}

func (p *Packet) push(bits ...rune) {
	p.bits = append(p.bits[:p.iter-1], append(bits, p.bits[p.iter:]...)...)
	p.iter--
}