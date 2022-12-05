package day4

import (
	"adventofcode/m/v2/util"
	"fmt"
	"regexp"
	"strconv"
)

func Day4(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Completely contained: %d\n", solveA(inputFile, incCompleteOverlap))
	} else {
		fmt.Printf("Partially contained: %d\n", solveA(inputFile, incPartialOverlap))
	}
}

func solveA(inputFile string, inc func(int, int, int, int) int) int {
	
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls) 
	
	re1 := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)
	var a,b,c,d, sum int
	sum = 0
	for ok {
		res := re1.FindStringSubmatch(line)

		a, _ = strconv.Atoi(res[1])
		b, _ = strconv.Atoi(res[2])
		c, _ = strconv.Atoi(res[3])
		d, _ = strconv.Atoi(res[4])

		// a-b contained in c-d if a >= c, b <= d
		// c-d contained in a-b if b >= a, d <= b 
		sum += inc(a,b,c,d)
		

		line, ok = util.Read(ls)
	}
	return sum
}

func incCompleteOverlap(a int, b int, c int, d int) int {
	if ( a >= c && b <= d ) || ( c >= a && d <= b ) {
		return 1
	}
	return 0
}

func incPartialOverlap(a int, b int, c int, d int) int {
	// a-b, c-d partially overlap if
	// a <= c <= b, a <= d <= b, c <= a <= d, c <= b <= d
	if ( a <= c && c <= b ) || ( a <= d && d <= b) || ( c <= a && a <= d) || ( c <= b && b <= d) {
		return 1 
	}
	return 0
}
