package main

import "fmt"

const testInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func main() {

	fmt.Printf("test part1: %d\n", part1(testInput))
	fmt.Printf("part1: %d\n", part1(input))

	fmt.Printf("test part2: %d\n", part2(testInput))
	fmt.Printf("part2: %d\n", part2(input))

}