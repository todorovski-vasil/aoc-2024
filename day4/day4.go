package main

import "fmt"

const testInput string = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func main() {
	fmt.Printf("test part1: %d\n", part1(testInput))
	fmt.Printf("part1: %d\n", part1(input))

	fmt.Printf("test part2: %d\n", part2(testInput))
	fmt.Printf("part2: %d\n", part2(input))
}