package util

import (
	"bufio"
	"fmt"
	"os"
)


func LineScanner(path string) *bufio.Scanner {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Can't open file:", os.Args[1])
		panic(err)
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	return s
}

func RuneScanner(path string) *bufio.Scanner {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Can't open file:", os.Args[1])
		panic(err)
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)
	return s
}

func WordScanner(path string) *bufio.Scanner {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Can't open file:", os.Args[1])
		panic(err)
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)
	return s
}

func Read(scanner *bufio.Scanner) (string, bool) {
	hasNext := scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning word")
		panic(err)
	}

	return scanner.Text(), hasNext
}
