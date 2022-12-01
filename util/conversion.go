package util

import (
	"fmt"
	"strconv"
)

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil{
		fmt.Printf("Couldn't convert %s to int", s)
	}

	return i
}

func Bin2Dec(binary string) int {
	i, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		fmt.Println("Error converting binary to decimal ", err)
	}

	return int(i)
}

func Hex2Bits(hex string) []uint64 {
	val, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		fmt.Printf("Error converting hex=%s to bits: %s\n", hex, err)
	}

	bits := []uint64{}
	for i := 0; i < 4; i++ {
		bits = append([]uint64{val & 0x1}, bits...)
		val = val >> 1
	}
	return bits
}
