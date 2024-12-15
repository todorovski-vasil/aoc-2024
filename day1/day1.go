package main

import "fmt"

const testInput = `3   4
4   3
2   5
1   3
3   9
3   3`

func main() {
	fmt.Printf("test part1: %d\n", part1(testInput))
	fmt.Printf("test part2: %d\n", part2(testInput))

	fmt.Printf("part1: %d\n", part1(input))
	fmt.Printf("part2: %d\n", part2(input))
}
