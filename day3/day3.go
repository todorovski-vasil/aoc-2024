package main

import "fmt"

const testInput = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func main() {
	fmt.Printf("test part1: %d\n", part1(testInput))
	fmt.Printf("part1: %d\n", part1(input))

	fmt.Printf("test part2: %d\n", part2(testInput))
	fmt.Printf("part2: %d\n", part2(input))
}