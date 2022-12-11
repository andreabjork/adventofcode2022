package day11

import (
	"adventofcode/m/v2/util"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Day11(inputFile string, part int) {
	monkeys, pc := makeMonkeys(inputFile)
	if part == 0 {
		fmt.Printf("Level of monkey business 20 rounds: %d\n", MonkeyBusiness(monkeys, 20, pc.divideByThree))
	} else if part == 1 {
	    fmt.Printf("Level of monkey business 1 round: %d\n", MonkeyBusiness(monkeys,20, pc.none))
	}
}

func (pc *PrimeCalculator) divideByThree(pf *PrimeFactor) *PrimeFactor {
	return pc.PF(pf.value()/3)
}

func (pc *PrimeCalculator) none(pf *PrimeFactor) *PrimeFactor {
	return pf
}

func MonkeyBusiness(monkeys []*Monkey, rounds int, manageWorry func(pf *PrimeFactor) *PrimeFactor) int {
	inspections := make([]int, len(monkeys))
	//showRound(monkeys)
	for r := 0; r < rounds; r++{
		for m := 0; m < len(monkeys); m++ {
			inspections[m] += len(monkeys[m].items)
			monkey := monkeys[m]
			var throwTo int
			hasItem, worry := monkey.pop()
			for hasItem {
				worry = monkey.op(worry)
				worry = manageWorry(worry)
				throwTo = monkey.modOp(worry)
				monkeys[throwTo].push(worry)
				hasItem, worry = monkey.pop()
			}
		}
		//showRound(monkeys)
	}
	for i := 0; i < len(inspections); i++ {
		fmt.Printf("Monkey %d inspected items %d times.\n", i, inspections[i])
	}
	sort.Ints(inspections)
	return inspections[len(inspections)-1]*inspections[len(inspections)-2]
}

func showRound(monkeys []*Monkey) {
	for i := 0; i < len(monkeys); i++{
		fmt.Printf("Monkey %d: ", i)
		for j := 0; j < len(monkeys[i].items); j++ {
			fmt.Printf("%d ", monkeys[i].items[j].value())
		}
		fmt.Println()
	}
}

// ================
// PRIME CALCULATOR
// ================
type PrimeCalculator struct {
	primes 	[]int
}

func (pc *PrimeCalculator) PF(x int) *PrimeFactor {
	pf := &PrimeFactor{map[int]int{}}
	for i := 1; i < len(pc.primes); i++ {
		for x % pc.primes[i] == 0 && x != 1 {
			x = x / pc.primes[i]
			if _, ok := pf.factors[pc.primes[i]]; ! ok {
				pf.factors[pc.primes[i]] = 1
			} else {
				pf.factors[pc.primes[i]]++
			}
		}
		if x == 1 {
			break
		}
	}
	return pf
}

func (pc *PrimeCalculator) add(pf, pg *PrimeFactor) *PrimeFactor {
	p := &PrimeFactor{map[int]int{}}
	pfRemainder:= 1
	pgRemainder:= 1
	for prime, c := range pg.factors {
		if cc, ok := pf.factors[prime]; ok {
			if c == cc {
				p.factors[prime] = c
			} else if c > cc {
				p.factors[prime] = cc
				pgRemainder *= util.Pow(prime,c-cc)
			} else {
				p.factors[prime] = c
				pfRemainder *= util.Pow(prime, cc-c)
			}
		} else {
			// pg prime not in pf
			pgRemainder *= util.Pow(prime, c)
		}
	}
	for prime, c := range pf.factors {
		if _, ok := pg.factors[prime]; !ok {
			// pg prime not in pf
			pfRemainder *= util.Pow(prime, c)
		}
	}
	pp := pc.multiply(p, pc.PF(pfRemainder + pgRemainder))
	if pp.value() != pf.value() + pg.value() {
		fmt.Println("WARNING: ADDITION ERROR")
		fmt.Printf("Wanted: %d+%d = %d, got %d\n", pf.value(), pg.value(), pf.value()+pg.value(), p.value())
	}
	return pp
}

func (pc *PrimeCalculator) multiply(pf, pg *PrimeFactor) *PrimeFactor {
	p := pf.copy()
	for prime, count := range pg.factors {
		if _, ok := p.factors[prime]; !ok {
			p.factors[prime] = count
		} else {
			p.factors[prime] += count
		}
	}
	if p.value() != pf.value()*pg.value() {
		fmt.Println("WARNING: MULTIPLICATION ERROR")
		fmt.Printf("Wanted: %d*%d = %d, got %d\n", pf.value(), pg.value(), pf.value()*pg.value(), p.value())
	}
	return p
}

func NewCalculator(n int) *PrimeCalculator {
	primes := []int{1}
	for i := 1; i <= n; i++ {
		x := i
		for p := 1; p < len(primes); p++ {
			for x % primes[p] == 0  && x != 1 {
				x = x / primes[p]
			}
		}
		// the remainder is a prime
		if x != 1 {
			primes = append(primes, x)
		}
	}
	return &PrimeCalculator{primes}
}

// =============
// PRIME FACTORS
// =============
type PrimeFactor struct {
	factors 	map[int]int // factors[i] = j means prime i occurs j times in this number
}

func (pf *PrimeFactor) value() int {
	val := 1
	for prime, count := range pf.factors {
		val *= util.Pow(prime, count)
	}
	return val
}

func (pf *PrimeFactor) print() {
	val := 1
	for prime, count := range pf.factors {
		fmt.Printf("%d(x%d) ", prime, count)
		val *= util.Pow(prime, count)
	}
	fmt.Printf("\nVal: %d\n\n", val)
}

func (pf *PrimeFactor) copy() *PrimeFactor {
	factors := map[int]int{}
	for prime, count := range pf.factors {
		factors[prime] = count
	}
	return &PrimeFactor{factors}
}

func (pf *PrimeFactor) moduloZero(prime int) bool {
	val, ok := pf.factors[prime]
	if ok != (pf.value() % prime == 0) {
		fmt.Println("WARNING: MODULO CALCULATION ERROR")
	}
	return ok && val >= 1
}

// =======
// MONKEY
// =======
type Monkey struct {
	items			[]*PrimeFactor
	op 				func(*PrimeFactor) *PrimeFactor
	modOp 			func(*PrimeFactor) int
}

func (m *Monkey) push(x *PrimeFactor) {
	m.items = append(m.items, x)
}

func (m *Monkey) pop() (bool, *PrimeFactor) {
	if len(m.items) == 0 {
		return false, nil
	}
	head := m.items[0]
	m.items = m.items[1:]
	return true, head
}

// =======
// PARSING
// =======
func makeMonkeys(inputFile string) ([]*Monkey, *PrimeCalculator) {
	txt, _ := os.ReadFile(inputFile)
	re := regexp.MustCompile(
		`Monkey (\d+):
  Starting items: ([\d+|,|\s]*)
  Operation: new = old (\+|\*) ([A-Za-z0-9]*)
  Test: divisible by (\d+)
    If true: throw to monkey (\d+)
    If false: throw to monkey (\d+)`)


	pc := NewCalculator(1000000)
	matches := re.FindAllStringSubmatch(string(txt), 10)
	monkeys := []*Monkey{}
	var itemWorries []string
	var operator, arg string
	for _, m := range matches {
		// Make starting items, representing their value with PrimeFactors
		itemWorries = strings.Split(m[2], ", ")
		items := make([]*PrimeFactor, len(itemWorries))
		for i := 0; i < len(items); i++ {
			x, _ := strconv.Atoi(itemWorries[i])
			items[i] = pc.PF(x)
		}

		// Set monkey operation
		operator = m[3]
		arg = m[4]
		var op func(*PrimeFactor) *PrimeFactor
		var y int
		if arg == "old" {
			switch operator {
			case "+":
				op = func(x *PrimeFactor) *PrimeFactor { return pc.add(x, x)}
			case "*":
				op = func(x *PrimeFactor) *PrimeFactor { return pc.multiply(x, x) }
			}
		} else {
			y, _ = strconv.Atoi(arg)
			switch operator {
			case "+":
				op = func(x *PrimeFactor) *PrimeFactor { return pc.add(x, pc.PF(y)) }
			case "*":
				op = func(x *PrimeFactor) *PrimeFactor { return pc.multiply(x, pc.PF(y)) }
			}
		}

		// Set monkey modulo test
		test, _ := strconv.Atoi(m[5])
		iftrue, _ := strconv.Atoi(m[6])
		iffalse, _ := strconv.Atoi(m[7])
		modOp := func(x *PrimeFactor) int {
			if x.moduloZero(test) {
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

	return monkeys, pc
}