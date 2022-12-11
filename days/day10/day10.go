package day10

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
)

func Day10(inputFile string, part int) {
	cpu := run(inputFile)
	if part == 0 {
		fmt.Printf("Sum signals: %d\n", cpu.regSum([]int{20,60,100,140,180,220}) )
	} else {
		fmt.Println("Screen image: ")
		cpu.image()
	}
}

func run(inputFile string) *CPU {
	ls := util.LineScanner(inputFile)

	cpu := &CPU{[]int{1}}
	line, ok := util.Read(ls)
	for ok {
		parts := strings.Split(line, " ")
		op := parts[0]
		if op == "noop" {
			cpu.noop()
		} else {
			x, _ := strconv.Atoi(parts[1])
			cpu.addx(x)
		}
		line, ok = util.Read(ls)
	}
	return cpu
}

type CPU struct {
	hist 	[]int
}

func (c *CPU) image() {
	for cyc := 1; cyc < len(c.hist); cyc++ {
		// Draw # if register-1 <= cyc <= register+1
		if c.reg(cyc)-1 <= (cyc-1)%40 && c.reg(cyc)+1 >= (cyc-1)%40 {
			fmt.Printf("#")
		} else if cyc %40 == 0 {
			fmt.Println()
		} else {
			fmt.Printf(".")
		}
	}
}

func (c *CPU) regSum(cycles []int) int {
	sum := 0
	for _, cyc := range cycles {
		sum += cyc*c.reg(cyc)
	}
	return sum
}

func (c *CPU) reg(cycle int) int {
	// Returning value during cycle, not after
	return c.hist[cycle-1]
}

func (c *CPU) noop() {
	c.hist = append(c.hist, c.hist[len(c.hist)-1])
}

func (c *CPU) addx(x int) {
	for i := 0; i < 1 ; i++{
		c.noop()
	}
	c.hist = append(c.hist, c.hist[len(c.hist)-1]+x)
}

