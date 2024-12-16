package main

import "fmt"

const testInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func main() {
	fmt.Printf("test part1: %d\n", part1(testInput))
	fmt.Printf("part1: %d\n", part1(input))

	fmt.Printf("test part2: %d\n", part2(testInput))
	fmt.Printf("part2: %d\n", part2(input))
}
