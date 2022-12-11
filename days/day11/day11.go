package day11

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Day11(inputFile string, part int) {
	monkeys := makeMonkeys(inputFile)
	if part == 0 {
		fmt.Printf("Level of monkey business 20 rounds: %d\n", MonkeyBusiness(monkeys, 20))
	} else {
		fmt.Printf("Level of monkey business 10k rounds: %d\n", MonkeyBusiness(monkeys, 10000))
	}
}

func MonkeyBusiness(monkeys []*Monkey, rounds int) int {
	inspections := make([]int, len(monkeys))

	//showRound(monkeys)
	for r := 0; r < rounds; r++{
		for m := 0; m < len(monkeys); m++ {
			inspections[m] += len(monkeys[m].items)
			inspectItems(m, monkeys)
		}
		//showRound(monkeys)
	}

	sort.Ints(inspections)
	return inspections[len(inspections)-1]*inspections[len(inspections)-2]
}

func showRound(monkeys []*Monkey) {
	for i := 0; i < len(monkeys); i++{
		fmt.Printf("Monkey %d: %+v\n", i, monkeys[i].items)
	}
}

type Monkey struct {
	items			[]int
	op 				func(int) int
	modOp 			func(int) int
}

func inspectItems(n int, monkeys []*Monkey) {
	m := monkeys[n]
	var throwTo int
	hasItem, worry := m.pop()
	for hasItem {
		worry = m.op(worry)/3
		throwTo = m.modOp(worry)
		monkeys[throwTo].push(worry)

		hasItem, worry = m.pop()
	}
}

func (m *Monkey) push(x int) {
	m.items = append(m.items, x)
}

func (m *Monkey) pop() (bool, int) {
	if len(m.items) == 0 {
		return false, 0
	}
	head := m.items[0]
	m.items = m.items[1:]
	return true, head
}

func makeMonkeys(inputFile string) []*Monkey {
	txt, _ := os.ReadFile(inputFile)
	re := regexp.MustCompile(
`Monkey (\d+):
  Starting items: ([\d+|,|\s]*)
  Operation: new = old (\+|-|\*|\/) ([A-Za-z0-9]*)
  Test: divisible by (\d+)
    If true: throw to monkey (\d+)
    If false: throw to monkey (\d+)`)

	matches := re.FindAllStringSubmatch(string(txt), 10)

	monkeys := []*Monkey{}
	var itemWorries []string
	var operator, arg string
	for _, m := range matches {
		// Make starting items
		itemWorries = strings.Split(m[2], ", ")
		items := make([]int, len(itemWorries))
		for i := 0; i < len(items); i++ {
			items[i], _ = strconv.Atoi(itemWorries[i])
		}

		// Set monkey operation
		operator = m[3]
		arg = m[4]
		var op func(int) int
		var y int
		if arg == "old" {
			switch operator {
			case "+":
				op = func(x int) int { return x+x }
			case "-":
				op = func(x int) int { return x-x }
			case "*":
				op = func(x int) int { return x*x }
			case "/":
				op = func(x int) int { return x/x }
			}
		} else {
			y, _ = strconv.Atoi(arg)
			switch operator {
			case "+":
				op = func(x int) int { return x+y }
			case "-":
				op = func(x int) int { return x-y }
			case "*":
				op = func(x int) int { return x*y }
			case "/":
				op = func(x int) int { return x/y }
			}
		}

		// Set monkey modulo test
		test, _ := strconv.Atoi(m[5])
		iftrue, _ := strconv.Atoi(m[6])
		iffalse, _ := strconv.Atoi(m[7])
		modOp := func(x int) int {
			if x % test == 0 {
				return iftrue
			} else {
				return iffalse
			}
		}

		monkeys = append(monkeys,
			&Monkey{
				items,
				op,
				modOp,
				},
			)
	}

	return monkeys
}