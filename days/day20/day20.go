package day20

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
)

func Day20(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Grove coordinate sum: %d\n", solve(inputFile))
	} else {
		fmt.Println("Not implmenented.")
	}
}

func solve(inputFile string) int {
	ls := util.LineScanner(inputFile)

	arr := []*Ele{}
	order := []*Ele{}
	line, ok := util.Read(ls)
	for ok {
		num, _ := strconv.Atoi(line)
		e := &Ele{num, 0}
		arr = append(arr, e)
		order = append(order, e)
		line, ok = util.Read(ls)
	}

	//fmt.Println("Len array: ", len(arr))
	//arr = applyKey(arr, 811589153)
	arr = decrypt(arr)
	zIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i].val == 0 {
			zIndex = i
			fmt.Println("FOUNDET", zIndex)
			break
		}
	}
	//fmt.Println(len(arr))
	//fmt.Println(arr[(zIndex+1000)%(len(arr)-1)].val)
	//fmt.Println(arr[(zIndex+2000)%(len(arr)-1)].val)
	//fmt.Println(arr[(zIndex+3000)%(len(arr)-1)].val)
	//print(arr)
	return arr[(zIndex+1000)%len(arr)].val + arr[(zIndex+2000)%len(arr)].val + arr[(zIndex+3000)%len(arr)].val

}

type Ele struct {
	val   int
	it    int
}

func decrypt(arr []*Ele) []*Ele {
	//fmt.Printf("%+v\n", arr)
	count := 0
	i := 0
	k := 0
	for count < len(arr) {
		fmt.Println(i)
		//fmt.Println(count)
		arr, k = move(arr, i)
		if k == -1 {
			i++
		} else {
			count++
			i += k
		}
	}
	return arr
}

func move(arr []*Ele, i int) ([]*Ele, int) {
	add := (i+arr[i].val)%(len(arr)-1)
	for add < 0 {
		add = (add+len(arr)-1)%(len(arr)-1)
	}
	var first, second []*Ele
	x := arr[i]
	if x.it != 0 {
		return arr, -1
	}
	k := -1
	if i <= add {
		first = append([]*Ele{}, arr[:i]...)
		first = append(first, arr[i+1:add+1]...)
		second = append([]*Ele{x}, arr[add+1:]...)
		x.it++
		k = 0
	} else {
		first = append([]*Ele{}, arr[:add]...)
		second = append([]*Ele{x}, arr[add:i]...)
		second = append(second, arr[i+1:]...)
		x.it++
		k = 1
	}

	arr = append(first, second...)
	//print(arr)
	return arr, k
}

func print(arr []*Ele) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i].val)
	}
	fmt.Println("")
}