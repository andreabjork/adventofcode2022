package main

import (
	"adventofcode/m/v2/days/day1"
	"adventofcode/m/v2/days/day2"
	"adventofcode/m/v2/days/day20"
	"adventofcode/m/v2/days/day3"
	"adventofcode/m/v2/days/day4"
	"adventofcode/m/v2/days/day5"
	"adventofcode/m/v2/days/day6"
	"adventofcode/m/v2/days/day7"
	"adventofcode/m/v2/days/day8"
	"adventofcode/m/v2/days/day9"
	"adventofcode/m/v2/days/day10"
	"adventofcode/m/v2/days/day11"
	"adventofcode/m/v2/days/day12"
	"adventofcode/m/v2/days/day13"
	"adventofcode/m/v2/days/day14"
	"adventofcode/m/v2/days/day15"
	"adventofcode/m/v2/days/day16"
	"adventofcode/m/v2/days/day17"
	"adventofcode/m/v2/days/day18"
	"adventofcode/m/v2/days/day19"
	"adventofcode/m/v2/days/day21"
	"adventofcode/m/v2/days/day22"
	"adventofcode/m/v2/days/day23"
	"adventofcode/m/v2/days/day24"
	"adventofcode/m/v2/days/day25"
	"adventofcode/m/v2/util"
	"fmt"
	"os"
	"time"
)

var dayfuncs = map[int]interface{}{
	1:  day1.Day1,
	2:  day2.Day2,
	3:  day3.Day3,
	4:  day4.Day4,
	5:  day5.Day5,
	6:  day6.Day6,
	7:  day7.Day7,
	8:  day8.Day8,
	9:  day9.Day9,
	10: day10.Day10,
	11: day11.Day11,
	12: day12.Day12,
	13: day13.Day13,
	14: day14.Day14,
	15: day15.Day15,
	16: day16.Day16,
	17: day17.Day17,
	18: day18.Day18,
	19: day19.Day19,
	20: day20.Day20,
	21: day21.Day21,
	22: day22.Day22,
	23: day23.Day23,
	24: day24.Day24,
	25: day25.Day25,
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Run as:\n\tgo run main.go [day] [0|1] [0-9:optional,picks 1.txt by default]")
		return
	}

	day := util.ToInt(os.Args[1])
	part := util.ToInt(os.Args[2])

	var input int
	if len(os.Args) > 3 {
		input = util.ToInt(os.Args[3])
	} else {
		input = 0
	}
	inputFile := fmt.Sprintf("days/day%d/input/%d.txt", day, input)

	runDay(day, part, inputFile)
}

func runDay(day int, part int, input string) {
	fmt.Printf("Day %d:%d < %s\n", day, part, input)
	start := time.Now()
	dayfuncs[day].(func(string, int))(input, part)
	elapsed := time.Now().Sub(start)
	fmt.Printf("Time: %v\n\n", elapsed)
}
