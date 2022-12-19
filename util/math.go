package util

func Max(x int, y int) int {
	if x >= y {
	return x
	}
	return y
}

func Min(x int, y int) int {
	if x <= y {
	return x
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sum(arr []int) int {
	sum := 0
	for _, a := range arr {
		sum += a
	}
	return sum
}

func Pow(x int, n int) int {
	val := 1
	for i := 0; i < n; i++ {
		val *= x
	}
	return val
}