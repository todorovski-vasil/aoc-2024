package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func part1(input string) int {
	var list1, list2 []int
	for _, line := range strings.Split(input, "\n") {
		var a, b int
		fmt.Sscanf(line, "%d %d", &a, &b)
		list1 = append(list1, a)
		list2 = append(list2, b)
	}
	sort.Ints(list1)
	sort.Ints(list2)
	var sum int

	for i := 0; i < len(list1); i++ {
		sum += int(math.Abs(float64(list1[i] - list2[i])))
	}
	return sum
}
