package day3

import (
  "adventofcode/m/v2/util"
	"fmt"
)

func Day3(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Sum of priorities: %d\n", solveA(inputFile))
	} else {
		fmt.Printf("Sum of priorities of all groups: %d\n", solveB(inputFile))
	}
}

func solveA(inputFile string) int { 
  // rucksack items a..z A..Z with priority 1..26 27..52  
  ls := util.LineScanner(inputFile)
  items := makeItems()

  prioritySum := 0
  line, ok := util.Read(ls)
  for ok {
    rucksackFront := map[rune]bool{}
    rucksackBack := map[rune]bool{}
    packedItems := []rune(line)
    N := len(packedItems)  

    for _, frontItem := range packedItems[:N/2] { 
      rucksackFront[frontItem] = true
    }

    for _, backItem := range packedItems[N/2:] { 
      if rucksackFront[backItem] && !rucksackBack[backItem] {
        rucksackBack[backItem] = true
        prioritySum += items[backItem]
      }
    }

    line, ok = util.Read(ls)
  }

  return prioritySum
}

func solveB(inputFile string) int {
  ls := util.LineScanner(inputFile)
  items := makeItems()

  line, ok := util.Read(ls)

  elves := 0
  prioritySum := 0
  // inRucksacks[r] = n says r is in n backpacks in this group
  inRucksacks := map[rune]int{}
  for ok {
    rucksack := map[rune]bool{}
    packedItems := []rune(line)
    for _, item := range packedItems { 
      if ! rucksack[item] {
        inRucksacks[item]++
        // Set priority for the group
        if inRucksacks[item] == 3 {
          fmt.Println("Found common item!", string(item))
          prioritySum += items[item]
        }
        rucksack[item] = true
      }
    }

    elves++
    line, ok = util.Read(ls)
    if elves == 3 {
      // Processing new group of elves 
      inRucksacks = map[rune]int{}
      elves = 0
    }
  }

  return prioritySum
}

func makeItems() map[rune]int {
  priority := 1 
  items := map[rune]int{}
  
  for r := 'a'; r <= 'z'; r++ {
    items[r] = priority
    priority++
  }
  
  for r := 'A'; r <= 'Z'; r++ {
    items[r] = priority
    priority++
  }

   return items
}
