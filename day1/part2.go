package main

import (
	"fmt"
	"strings"
)

func part2(input string) int {
	var list1, list2 []int
	for _, line := range strings.Split(input, "\n") {
		var a, b int
		fmt.Sscanf(line, "%d %d", &a, &b)
		list1 = append(list1, a)
		list2 = append(list2, b)
	}
	var sum int

	dict2 := make(map[int]int)
	for _, num := range list2 {
		dict2[num] = dict2[num] + 1
	}

	for i := 0; i < len(list1); i++ {

		sum += list1[i] * dict2[list1[i]]
	}
	return sum
}
