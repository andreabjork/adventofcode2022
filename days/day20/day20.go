package day20

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"time"
)

func Day20(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Grove coordinate sum: %d\n", solve(inputFile, 1, 1))
	} else {
		fmt.Printf("Grove coordinate sum: %d\n", solve(inputFile, 10, 811589153))
	}
}

func solve(inputFile string, times int, key int) int {
	ls := util.LineScanner(inputFile)

	arr := []*Ele{}
	order := []*Ele{}
	line, ok := util.Read(ls)
	it := 0
	for ok {
		num, _ := strconv.Atoi(line)
		e := &Ele{num*key, it}
		it++
		arr = append(arr, e)
		order = append(order, e)
		line, ok = util.Read(ls)
	}

	arr = decrypt(arr, times)
	//fmt.Println("After:")
	//print(arr)
	zIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i].val == 0 {
			zIndex = i
			//fmt.Println("Found it", zIndex)
			break
		}
	}

	//fmt.Println("z + 1000", zIndex+1000)
	//fmt.Println("z + 1000 mod len ", (zIndex+1000)%len(arr))
	//fmt.Println("val ", arr[(zIndex+1000)%len(arr)])
	//fmt.Println(arr[(zIndex+1000)%len(arr)].val)
	//fmt.Println(arr[(zIndex+2000)%len(arr)].val)
	//fmt.Println(arr[(zIndex+3000)%len(arr)].val)
	//fmt.Println(arr[(zIndex+1000)%(len(arr)-1)].val)
	//fmt.Println(arr[(zIndex+2000)%(len(arr)-1)].val)
	//fmt.Println(arr[(zIndex+3000)%(len(arr)-1)].val)
	return arr[(zIndex+1000)%len(arr)].val + arr[(zIndex+2000)%len(arr)].val + arr[(zIndex+3000)%len(arr)].val

}

type Ele struct {
	val   int
	it    int
}

func decrypt(arr []*Ele, times int) []*Ele {
	count := 0
	i := 0
	//print(arr)
	for count < times*len(arr) {
		arr = move(arr, i)
		i++
		count++
		if count % 10 == 0 {
			fmt.Printf("%s: %d / %d\n", time.Now(), count, times*len(arr))
		}
		//print(arr)
	}
	return arr
}

func move(arr []*Ele, i int) []*Ele {
	var removeFrom int
	var insertAt int
	var x *Ele
	for j := 0; j < len(arr); j++ {
		if arr[j].it == i {
		   removeFrom = j
		   x = arr[j]
		   x.it = x.it + len(arr)
		   break
		}
	}

	//fmt.Println("before")
	//print(arr)
	arr = remove(arr, removeFrom)
	//fmt.Println("remove")
	//print(arr)
	insertAt = removeFrom+x.val
	for insertAt <= 0 {
		insertAt += len(arr)
	}
	insertAt = insertAt%len(arr)

	// This is complete defeat after not getting append to behave as expected
	arr = insert(arr, x, insertAt)
	//fmt.Println("iterator before insert", x.it)
	//fmt.Println("insert at ", x.val, insertAt)
	//print(arr)
	return arr
}

func remove(arr []*Ele, idx int) []*Ele {
	return append(arr[:idx], arr[idx+1:]...)
}

func insert(arr []*Ele, ele *Ele, idx int) []*Ele {
	last := len(arr) - 1
	arr = append(arr, arr[last])           // Step 1
	copy(arr[idx+1:], arr[idx:last]) // Step 2
	arr[idx] = ele
	return arr
}

func print(arr []*Ele) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i].val)
	}
	fmt.Println("")
}