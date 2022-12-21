package day21

import (
	"adventofcode/m/v2/util"
	"fmt"
	"regexp"
	"strconv"
)

func Day21(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Root: %d\n", solveA(inputFile))
	} else {
		solveB(inputFile)
	}
}

func solveA(inputFile string) int {
	re := regexp.MustCompile(`([a-z]{4}): (\d*[a-z]*)\s*([-+*/]*)\s*([a-z]{4})*`)

	monkeys := map[string]func() int{}
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)
	for ok {
		parts := re.FindStringSubmatch(line)
		op := ""
		if len(parts) > 3 {
			op = parts[3]
		}

		name := parts[1]
		switch op {
		case "-":
			monkeys[name] = func() int { return monkeys[parts[2]]() - monkeys[parts[4]]() }
		case "+":
			monkeys[name] = func() int { return monkeys[parts[2]]() + monkeys[parts[4]]() }
		case "*":
			monkeys[name] = func() int { return monkeys[parts[2]]() * monkeys[parts[4]]() }
		case "/":
			monkeys[name] = func() int { return monkeys[parts[2]]() / monkeys[parts[4]]() }
		default:
			num, _ := strconv.Atoi(parts[2])
			monkeys[name] = func() int { return num }
		}

		line, ok = util.Read(ls)
	}

	return monkeys["root"]()
}

func solveB(inputFile string) int {
	re := regexp.MustCompile(`([a-z]{4}): (\d*[a-z]*)\s*([-+*/]*)\s*([a-z]{4})*`)

	monkeys := map[string]func(bool, int) (bool, int){}
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)
	for ok {
		parts := re.FindStringSubmatch(line)
		op := ""
		if len(parts) > 3 {
			op = parts[3]
		}

		name := parts[1]
		if name == "root" {
			monkeys[name] = func(request bool, want int) (bool, int) {
				ok, x := monkeys[parts[2]](false, 0)
				var y int
				if ok {
					monkeys[parts[4]](true, x)
				} else {
					_, y = monkeys[parts[4]](false, 0)
					monkeys[parts[2]](true, y)
				}
				return true, x + y
			}
		} else {
			switch op {
			case "-":
				// want = x - y
				monkeys[name] = func(request bool, want int) (bool, int) {
					okX, x := monkeys[parts[2]](false, 0)
					okY, y := monkeys[parts[4]](false, 0)
					if okX && okY {
						return true, x - y
					} else if !request {
						return false, 0
					} else if okX {
						okY, y := monkeys[parts[4]](true, x-want)
						return okY, x-y
					} else {
						okX, x := monkeys[parts[2]](true, want+y)
						return okX, x-y
					}
				}
			case "+":
				// want = x + y
				monkeys[name]  = func(request bool, want int) (bool, int) {
					okX, x := monkeys[parts[2]](false, 0)
					okY, y := monkeys[parts[4]](false, 0)
					if okX && okY {
						return true, x + y
					} else if !request {
						return false, 0
					} else if okX {
						okY, y := monkeys[parts[4]](true, want-x)
						return okY, x+y
					} else {
						okX, x := monkeys[parts[2]](true, want-y)
						return okX, x+y
					}
				}
			case "*":
				// want = x * y
				monkeys[name] = func(request bool, want int) (bool, int) {
					okX, x := monkeys[parts[2]](false, 0)
					okY, y := monkeys[parts[4]](false, 0)
					var newWant int
					if okX && okY {
						return true, x * y
					} else if !request {
						return false, 0
					} else if okX {
						if x == 0 { newWant = 0 } else { newWant = want/x }
						okY, y := monkeys[parts[4]](true, newWant)
						return okY, x*y
					} else {
						if y == 0 { newWant = 0 } else { newWant = want/y }
						okX, x := monkeys[parts[2]](true, newWant)
						return okX, x*y
					}
				}
			case "/":
				// want = x / y
				monkeys[name] = func(request bool, want int) (bool, int) {
					okX, x := monkeys[parts[2]](false, 0)
					okY, y := monkeys[parts[4]](false, 0)
					if okX && okY {
						return true, x/y
					} else if !request {
						return false, 0
					} else if okX {
						var newWant int
						if want == 0 { newWant = 0 } else { newWant = x/want }
						okY, y := monkeys[parts[4]](true, newWant)
						return okY, x/y
					} else {
						okX, x = monkeys[parts[2]](true, want*y)
						return okX, x/y
					}
				}
			default:
				num, _ := strconv.Atoi(parts[2])
				monkeys[name] = func(request bool, want int) (bool, int) {
					return true, num
				}
			}
		}

		line, ok = util.Read(ls)
	}

	monkeys["humn"] = func(request bool, want int) (bool, int) {
		if request {
			fmt.Printf("Yell %d!\n", want)
			return true, want
		} else {
			return false, 0
		}
	}

	_, x := monkeys["root"](false, 0)
	return x
}