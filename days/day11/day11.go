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
	var rounds int
	var curbWorry bool
	if part == 0 {
		rounds = 20
		curbWorry = true
	} else if part == 1 {
		rounds = 10000
		curbWorry = false
	}
	monkeys := makeMonkeys(inputFile, curbWorry)
	fmt.Printf("Level of monkey business: %d\n", MonkeyBusiness(monkeys, rounds, curbWorry))
}

func MonkeyBusiness(monkeys []*Monkey, rounds int, curbWorry bool) int {
	inspections := make([]int, len(monkeys))
	for r := 0; r < rounds; r++{
		for m := 0; m < len(monkeys); m++ {
			inspections[m] += len(monkeys[m].items)
			monkey := monkeys[m]
			var throwTo int
			hasItem, item := monkey.pop()
			for hasItem {
				monkey.op(item)
				if curbWorry {
					item.value = item.value/3
					item.reconcile()
				}
				throwTo = monkey.modOp(item)
				monkeys[throwTo].push(item)
				hasItem, item = monkey.pop()
			}
		}
	}
	for i := 0; i < len(inspections); i++ {
		fmt.Printf("Monkey %d inspected items %d times.\n", i, inspections[i])
	}
	sort.Ints(inspections)
	return inspections[len(inspections)-1]*inspections[len(inspections)-2]
}

// =======
// MONKEYS
// =======
type Monkey struct {
	items			[]*Item
	mod 			int
	op 				func(*Item)
	modOp 			func(*Item) int
}

func (m *Monkey) push(x *Item) {
	m.items = append(m.items, x)
}

func (m *Monkey) pop() (bool, *Item) {
	if len(m.items) == 0 {
		return false, nil
	}
	head := m.items[0]
	m.items = m.items[1:]
	return true, head
}

// =====
// ITEMS
// =====
type Item struct {
	// Keys are all prime
	modulos		map[int]int
	value 		int
}

// Reconciles modulo map by the current value (pt 1)
func (i *Item) reconcile() {
	for m, _ := range i.modulos {
		i.modulos[m] = i.value % m
	}
}

func (i *Item) mod(m int) int {
	return  i.modulos[m]
}

func add(x int) func(*Item) {
	// (x + y) mod m == x mod m + y mod m
	return func (i *Item) {
		if i.value != -1 {
			i.value += x
		}
		for m, _ := range i.modulos {
			i.modulos[m] = (i.modulos[m] + x % m)%m
		}
	}
}

func double() func(*Item) {
	// (x + x) mod m = x mod m + x mod m
	return func(i *Item) {
		if i.value != -1 {
			i.value += i.value
		}
		for m, _ := range i.modulos {
			i.modulos[m] = (i.modulos[m] + i.modulos[m]) % m
		}
	}
}

func pow2() func(*Item) {
	return func(i *Item) {
		if i.value != -1 {
			i.value *= i.value
		}
		for m, _ := range i.modulos {
			i.modulos[m] = (i.modulos[m]*i.modulos[m])%m
		}
	}
}

func multiply(p int) func(*Item) {
	return func(i *Item) {
		if i.value != -1 {
			i.value *= p
		}
		for m, _ := range i.modulos {
			i.modulos[m] = (p*i.modulos[m])%m
		}
	}
}

// =============
// PARSING INPUT
// =============
func makeMonkeys(inputFile string, trackValues bool) []*Monkey {
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
	modulosToTrack := []int{}
	for _, m := range matches {
		// Make starting items
		itemWorries = strings.Split(m[2], ", ")

		// Set monkey operation
		operator = m[3]
		arg = m[4]
		var op func(*Item)
		var y int
		if arg == "old" {
			switch operator {
			case "+":
				op = double()
			case "*":
				op = pow2()
			}
		} else {
			y, _ = strconv.Atoi(arg)
			switch operator {
			case "+":
				op = add(y)
			case "*":
				op = multiply(y)
			}
		}

		// Set monkey modulo test
		prime, _ := strconv.Atoi(m[5])
		iftrue, _ := strconv.Atoi(m[6])
		iffalse, _ := strconv.Atoi(m[7])
		modOp := func(i *Item) int {
			if i.mod(prime) == 0 {
				return iftrue
			} else {
				return iffalse
			}
		}
		modulosToTrack = append(modulosToTrack, prime)

		items := make([]*Item, len(itemWorries))
		for i := 0; i < len(itemWorries); i++ {
			worry, _ := strconv.Atoi(itemWorries[i])
			items[i] = &Item{
				map[int]int{},
				worry,
			}
		}

		monkeys = append(monkeys,
			&Monkey{
				items,
				prime,
				op,
				modOp,
			},
		)
	}

	// Initialize all modulos
	for _, monkey := range monkeys {
		for _, item := range monkey.items {
			for _, modulo := range modulosToTrack {
				item.modulos[modulo] = item.value % modulo
			}
			if !trackValues {
				item.value = -1
			}
		}
	}

	return monkeys
}
